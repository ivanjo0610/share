// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dkim

import (
	"bytes"
	"fmt"

	"github.com/shuLhan/share/lib/dns"
)

var dnsClientPool *dns.UDPClientPool // nolint: gochecknoglobals

//
// Key represent a DKIM key record.
//
type Key struct {
	// REQUIRED fields.

	// Public contains public key data.
	// ("p=", base64, REQUIRED)
	Public []byte

	// RECOMMENDED fields.

	// Version of DKIM key record.
	// ("v=", plain-text, RECOMMENDED, default is "DKIM1")
	Version []byte

	// OPTIONAL fields.

	// Type of key.
	// ("k=", plain-text, OPTIONAL, default is "rsa").
	Type KeyType

	// HashAlgs contains list of hash algorithm that might be used.
	// ("h=", plain-text colon-separated,  OPTIONAL, defaults to allowing
	// all algorithms)
	HashAlgs []HashAlg

	// Notes that might be interest to human.
	// ("n=", qp-section, OPTIONAL, default is empty)
	Notes []byte

	// Services contains list of service types to which this record
	// applies.
	// ("s=", plain-text colon-separated, OPTIONAL, default is "*").
	Services [][]byte

	// Flags contains list of flags.
	// ("t=", plain-text colon-separated, OPTIONAL, default is no flags set)
	Flags []KeyFlag
}

//
// LookupKey DKIM (public) key using specific query method, domain name
// (SDID), and selector.
// The sdid MUST NOT be empty, but the selector MAY be empty.
//
func LookupKey(qmethod QueryMethod, sdid, selector []byte) (key *Key, err error) {
	if len(sdid) == 0 {
		return nil, fmt.Errorf("dkim: LookupKey: empty SDID")
	}
	if qmethod.Type == QueryTypeDNS {
		key, err = lookupDNS(qmethod.Option, sdid, selector)
	}
	return key, err
}

//
// ParseTXT parse DNS TXT resource record into Key.
//
func ParseTXT(txt []byte) (key *Key, err error) {
	p := newParser(txt)

	key = &Key{}
	for {
		tag, err := p.fetchTag()
		if err != nil {
			return nil, err
		}
		if tag == nil {
			break
		}
		err = key.set(tag)
		if err != nil {
			return nil, err
		}
	}

	return key, nil
}

func lookupDNS(opt QueryOption, sdid, selector []byte) (key *Key, err error) {
	if opt == QueryOptionTXT {
		key, err = lookupDNSTXT(sdid, selector)
	}
	return key, err
}

func lookupDNSTXT(sdid, selector []byte) (key *Key, err error) {
	if dnsClientPool == nil {
		err = newDNSClientPool()
		if err != nil {
			return nil, err
		}
	}

	dnsClient := dnsClientPool.Get()

	var bb bytes.Buffer
	if len(selector) > 0 {
		bb.Write(selector)
		bb.WriteByte('.')
	}
	bb.Write(dkimSubdomain)
	bb.WriteByte('.')
	bb.Write(sdid)

	qname := bb.Bytes()
	dnsMsg, err := dnsClient.Lookup(dns.QueryTypeTXT, dns.QueryClassIN, qname)
	if err != nil {
		dnsClientPool.Put(dnsClient)
		return nil, fmt.Errorf("dkim: LookupKey: " + err.Error())
	}
	if len(dnsMsg.Answer) == 0 {
		dnsClientPool.Put(dnsClient)
		return nil, fmt.Errorf("dkim: LookupKey: empty answer on '%s'", qname)
	}

	dnsClientPool.Put(dnsClient)

	answers := dnsMsg.FilterAnswers(dns.QueryTypeTXT)
	if len(answers) == 0 {
		return nil, fmt.Errorf("dkim: LookupKey: no TXT record on '%s'", qname)
	}
	if len(answers) != 1 {
		return nil, fmt.Errorf("dkim: LookupKey: multiple TXT record on '%s'", qname)
	}

	txt := answers[0].RData().([]byte)

	return ParseTXT(txt)
}

func newDNSClientPool() (err error) {
	var ns []string

	if len(DefaultNameServers) > 0 {
		ns = DefaultNameServers
	} else {
		ns = dns.GetSystemNameServers("")
		if len(ns) == 0 {
			ns = append(ns, "1.1.1.1")
		}
	}

	dnsClientPool, err = dns.NewUDPClientPool(ns)

	return err
}

//
// Bytes return text representation of Key.
//
func (key *Key) Bytes() []byte {
	var bb bytes.Buffer

	if len(key.Version) > 0 {
		bb.WriteString("v=")
		bb.Write(key.Version)
		bb.WriteByte(';')
	}
	if len(key.Public) > 0 {
		if bb.Len() > 0 {
			bb.WriteByte(' ')
		}
		bb.WriteString("p=")
		bb.Write(key.Public)
		bb.WriteByte(';')
	}

	bb.WriteString(" k=")
	bb.Write(keyTypeNames[key.Type])
	bb.WriteByte(';')

	if len(key.HashAlgs) > 0 {
		bb.WriteString(" h=")
		bb.Write(packHashAlgs(key.HashAlgs))
		bb.WriteByte(';')
	}
	if len(key.Notes) > 0 {
		bb.WriteString(" n=")
		bb.Write(key.Notes)
		bb.WriteByte(';')
	}
	if len(key.Services) > 0 {
		bb.WriteString(" s=")
		bb.Write(bytes.Join(key.Services, sepColon))
		bb.WriteByte(';')
	}
	if len(key.Flags) > 0 {
		bb.WriteString(" t=")
		bb.Write(packKeyFlags(key.Flags))
	}

	return bb.Bytes()
}

func (key *Key) set(t *tag) (err error) {
	if t == nil {
		return nil
	}
	switch t.key {
	case tagDNSPublicKey:
		key.Public = t.value

	case tagDNSVersion:
		key.Version = t.value

	case tagDNSHashAlgs:
		key.HashAlgs = unpackHashAlgs(t.value)

	case tagDNSKeyType:
		key.Type = parseKeyType(t.value)

	case tagDNSNotes:
		key.Notes = t.value

	case tagDNSServices:
		key.Services = bytes.Split(t.value, sepColon)

	case tagDNSFlags:
		key.Flags = unpackKeyFlags(t.value)
	}
	return err
}

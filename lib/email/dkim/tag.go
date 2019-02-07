// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dkim

import (
	"bytes"
	"fmt"

	libbytes "github.com/shuLhan/share/lib/bytes"
)

type tagKey int

//
// List of known tag keys.
//
const (
	tagUnknown tagKey = 0

	// Tags in DKIM-Signature field value, ordered by priority.

	// Required tags.
	tagVersion   tagKey = 1 << iota // v=
	tagAlg                          // a=
	tagSDID                         // d=
	tagSelector                     // s=
	tagHeaders                      // h=
	tagBodyHash                     // bh=
	tagSignature                    // b=
	// Recommended tags.
	tagCreatedAt // t=
	tagExpiredAt // x=
	// Optional tags.
	tagCanon          // c=
	tagPresentHeaders // z=
	tagAUID           // i=
	tagBodyLength     // l=
	tagQueryMethod    // q=

	// Tags in DNS TXT record.

	// Required tags.
	tagDNSPublicKey // p=
	// Optional tags.
	tagDNSKeyType // k=
	tagDNSNotes   // n=
)

//
// Mapping between tag in DKIM-Signature and tag in DKIM domain record,
// since both have the same text representation.
//
const (
	// Recommended tags.
	tagDNSVersion tagKey = tagVersion // v=
	// Optional tags.
	tagDNSHashAlgs = tagHeaders   // h=
	tagDNSServices = tagSelector  // s=
	tagDNSFlags    = tagCreatedAt // t=
)

//
// Mapping between tag key in numeric and their human readable form.
//
var tagKeys = map[tagKey][]byte{ // nolint: gochecknoglobals
	tagVersion:        []byte("v"),
	tagAlg:            []byte("a"),
	tagSDID:           []byte("d"),
	tagSelector:       []byte("s"),
	tagHeaders:        []byte("h"),
	tagBodyHash:       []byte("bh"),
	tagSignature:      []byte("b"),
	tagCreatedAt:      []byte("t"),
	tagExpiredAt:      []byte("x"),
	tagCanon:          []byte("c"),
	tagPresentHeaders: []byte("z"),
	tagAUID:           []byte("i"),
	tagBodyLength:     []byte("l"),
	tagQueryMethod:    []byte("q"),

	tagDNSPublicKey: []byte("p"),
	tagDNSKeyType:   []byte("k"),
	tagDNSNotes:     []byte("n"),
}

//
// tag represent tag's key and its value.
//
type tag struct {
	key   tagKey
	value []byte
}

//
// newTag create new tag only if key is valid: start with ALPHA and contains
// only ALPHA, DIGIT, or "_".
//
func newTag(key []byte) (t *tag, err error) {
	if len(key) == 0 {
		return nil, nil
	}
	if !libbytes.IsAlpha(key[0]) {
		return nil, fmt.Errorf("dkim: invalid key: '%s'", key)
	}
	for x := 0; x < len(key); x++ {
		if libbytes.IsAlnum(key[x]) || key[x] == '_' {
			continue
		}
		return nil, fmt.Errorf("dkim: invalid key: '%s'", key)
	}

	t = &tag{
		key: tagUnknown,
	}
	for k, v := range tagKeys {
		if bytes.Equal(v, key) {
			t.key = k
			break
		}
	}

	return t, nil
}

//
// setValue set the tag value only if val is valid, otherwise return an error.
//
func (t *tag) setValue(val []byte) (err error) {
	if len(val) == 0 {
		return nil
	}
	for x := 0; x < len(val); x++ {
		if libbytes.IsSpace(val[x]) {
			continue
		}
		if val[x] < 0x21 || val[x] == 0x3B || val[x] > 0x7E {
			return fmt.Errorf("dkim: invalid value: '%s'", val)
		}
	}

	t.value = val

	return nil
}

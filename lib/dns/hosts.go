package dns

import (
	"net"
	"runtime"

	libbytes "github.com/shuLhan/share/lib/bytes"
	libfile "github.com/shuLhan/share/lib/file"
	libnet "github.com/shuLhan/share/lib/net"
)

const (
	HostsFilePOSIX   = "/etc/hosts"
	HostsFileWindows = "C:\\Windows\\System32\\Drivers\\etc\\hosts"
	defaultTTL       = 604800 // 7 days
)

//
// GetSystemHosts return path to system hosts file.
//
func GetSystemHosts() string {
	if runtime.GOOS == "windows" {
		return HostsFileWindows
	}
	return HostsFilePOSIX
}

//
// HostsLoad parse the content of hosts file as packed DNS message.
// If path is empty, it will load from the system hosts file.
//
func HostsLoad(path string) (msgs []*Message, err error) {
	if len(path) == 0 {
		path = GetSystemHosts()
	}

	reader, err := libfile.NewReader(path)
	if err != nil {
		return
	}

	msgs = parse(reader)

	return
}

func newMessage(addr, hname []byte) *Message {
	if !libnet.IsHostnameValid(hname) {
		return nil
	}
	ip := net.ParseIP(string(addr))
	if ip == nil {
		return nil
	}

	qtype := QueryTypeA
	for x := 0; x < len(addr); x++ {
		if addr[x] == ':' {
			qtype = QueryTypeAAAA
			break
		}
	}

	libbytes.ToLower(&hname)
	rrName := make([]byte, len(hname))
	copy(rrName, hname)

	msg := &Message{
		Header: &SectionHeader{
			QDCount: 1,
			ANCount: 1,
		},
		Question: &SectionQuestion{
			Name:  hname,
			Type:  qtype,
			Class: QueryClassIN,
		},
		Answer: []*ResourceRecord{{
			Name:  rrName,
			Type:  qtype,
			Class: QueryClassIN,
			TTL:   defaultTTL,
			Text: &RDataText{
				v: addr,
			},
		}},
	}

	_, err := msg.MarshalBinary()
	if err != nil {
		return nil
	}

	return msg
}

//
// Fields of the entry are separated by any number of blanks and/or tab
// characters.
// Text from a "#" character until the end of the line is a comment, and is
// ignored.
// Host names may contain only alphanumeric characters, minus signs ("-"), and
// periods (".").
// They must begin with an alphabetic character and end with an alphanumeric
// character.
// Optional aliases provide for name changes, alternate spellings,
// shorter hostnames, or generic hostnames (for example, localhost). [1]
//
// [1] man 5 hosts
//
func parse(reader *libfile.Reader) (msgs []*Message) {
	var (
		seps  = []byte{'\t', '\v', ' '}
		terms = []byte{'\n', '\f', '#'}
	)
	for {
		c := reader.SkipSpace()
		if c == 0 {
			break
		}
		if c == '#' {
			reader.SkipUntilNewline()
			continue
		}

		addr, isTerm, c := reader.ReadUntil(seps, terms)
		if isTerm {
			if c == 0 {
				break
			}
			if c == '#' {
				reader.SkipUntilNewline()
			}
			continue
		}

		for {
			c := reader.SkipSpace()
			if c == 0 {
				break
			}
			if c == '#' {
				reader.SkipUntilNewline()
				break
			}
			hname, isTerm, c := reader.ReadUntil(seps, terms)
			if len(hname) > 0 {
				msg := newMessage(addr, hname)
				if msg != nil {
					msgs = append(msgs, msg)
				}
			}
			if isTerm {
				if c == 0 {
					break
				}
				if c == '#' {
					reader.SkipUntilNewline()
				}
				break
			}
		}
	}

	return
}

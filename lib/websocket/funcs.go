// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"bytes"
	"crypto/sha1" //nolint: gosec
	"encoding/base64"
	"encoding/binary"
	"math/rand"
	"time"

	"golang.org/x/sys/unix"
)

//
// Recv read all content from file descriptor into slice of bytes.
//
// On success it will return buffer from pool. Caller must put the buffer back
// to the pool.
//
// On fail it will return nil buffer and error.
//
func Recv(fd int) (packet []byte, err error) {
	bs := _bsPool.Get().(*[]byte)

	n, err := unix.Read(fd, *bs)
	if err != nil {
		_bsPool.Put(bs)
		return nil, err
	}
	if n == 0 {
		_bsPool.Put(bs)
		return nil, nil
	}

	bb := _bbPool.Get().(*bytes.Buffer)
	bb.Reset()

	for n == _maxBuffer {
		_, err = bb.Write((*bs)[:n])
		if err != nil {
			goto out
		}

		n, err = unix.Read(fd, *bs)
		if err != nil {
			goto out
		}

	}
	if n > 0 {
		_, err = bb.Write((*bs)[:n])
		if err != nil {
			goto out
		}
	}

out:
	if err == nil {
		packet = make([]byte, bb.Len())
		copy(packet, bb.Bytes())
	}

	_bbPool.Put(bb)
	_bsPool.Put(bs)

	return packet, err
}

//
// Send the packet through web socket file descriptor `fd`.
//
func Send(fd int, packet []byte) (err error) {
	_, err = unix.Write(fd, packet)

	return
}

//
// SendFrame by packing websocket frame first and write into it.
//
func SendFrame(fd int, f *Frame, randomMask bool) (err error) {
	if fd == 0 || f == nil {
		return
	}

	packet := f.Pack(randomMask)

	return Send(fd, packet)
}

func concatBytes(bs0 []byte, bs1 ...byte) (out []byte) {
	out = append(out, bs0...)
	out = append(out, bs1...)

	return
}

//
// generateHandshakeAccept generate server accept key by concatenating key,
// defined in step 4 in Section 4.2.2, with the string
// "258EAFA5-E914-47DA-95CA-C5AB0DC85B11", taking the SHA-1 hash of this
// concatenated value to obtain a 20-byte value and base64-encoding (see
// Section 4 of [RFC4648]) this 20-byte hash.
//
func generateHandshakeAccept(key []byte) string {
	key = append(key, "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"...)
	sum := sha1.Sum(key) // nolint: gosec
	return base64.StdEncoding.EncodeToString(sum[:])
}

//
// generateHandshakeKey randomly selected 16-byte value that has been
// base64-encoded (see Section 4 of [RFC4648]).
//
func generateHandshakeKey() (key []byte) {
	if _rng == nil {
		_rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	bkey := make([]byte, 16)

	binary.LittleEndian.PutUint64(bkey[0:8], _rng.Uint64())
	binary.LittleEndian.PutUint64(bkey[8:16], _rng.Uint64())

	key = make([]byte, base64.StdEncoding.EncodedLen(len(bkey)))
	base64.StdEncoding.Encode(key, bkey)

	return
}
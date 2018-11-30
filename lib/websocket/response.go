// Copyright 2018, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"sync"
)

var ( // nolint: gochecknoglobals
	_resPool = sync.Pool{
		New: func() interface{} {
			return new(Response)
		},
	}
)

//
// Response contains the data that server send to client as a reply from
// Request or as broadcast from client subscription.
//
// If response type is a reply from Request, the ID from Request will be
// copied to the Response, and `code` and `message` field values are equal
// with HTTP response code and message.
//
// Example of response format for replying request,
//
// 	{
// 		id: 1512459721269,
// 		code:  200,
// 		message: "",
// 		body: ""
// 	}
//
// If response type is broadcast the ID and code MUST be 0, and the `message`
// field will contain the name of subscription. For example, when recipient of
// message read the message, server will publish a notification response as,
//
// 	{
// 		id: 0,
// 		code:  0,
// 		message: "message.read",
// 		body: "{ \"id\": ... }"
// 	}
//
type Response struct {
	ID      uint64 `json:"id"`
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Body    string `json:"body"`
}

//
// Reset all field's value to zero or empty.
//
func (res *Response) Reset() {
	res.ID = 0
	res.Code = 0
	res.Message = ""
	res.Body = ""
}

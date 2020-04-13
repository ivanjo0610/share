// Copyright 2020, Shulhan <m.shulhan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bot

//
// PassportData contains information about Telegram Passport data shared with
// the bot by the user.
type PassportData struct {
	// Array with information about documents and other Telegram Passport
	// elements that was shared with the bot.
	data []EncryptedPassportElement

	// Encrypted credentials required to decrypt the data.
	Credentials EncryptedCredentials
}
// Copyright 2012, Jeramey Crawford <jeramey@antihe.ro>
// Copyright 2013, Jonas mg
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package sha512_test

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/metrumresearchgroup/sha512"
)

var _Salt = &Salt{
	MagicPrefix:   []byte("$foo$"),
	SaltLenMin:    1,
	SaltLenMax:    8,
	RoundsMin:     1000,
	RoundsMax:     999999999,
	RoundsDefault: 5000,
}

func TestGenerateSaltWRounds(t *testing.T) {
	const rounds = 5001
	salt, err := _Salt.GenerateWRounds(_Salt.SaltLenMax, rounds)
	if err != nil {
		t.Fatal(err)
	}
	if salt == nil {
		t.Errorf("salt should not be nil")
	}

	expectedPrefix := string(_Salt.MagicPrefix) + "rounds=" + strconv.Itoa(rounds) + "$"
	if !strings.HasPrefix(string(salt), expectedPrefix) {
		t.Errorf("salt '%s' should start with prefix '%s' but didn't", salt, expectedPrefix)
	}
}

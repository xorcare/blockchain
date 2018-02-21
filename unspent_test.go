// Copyright 2017 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import "testing"

func TestGetUnspent(t *testing.T) {
	addresses := []string{
		"1J6PYEzr4CUoGbnXrELyHszoTSz3wCsCaj",
		"12cbQLTFMXRnSzktFkuoG3eHoMeFtpTu3S",
		"15yN7NPEpu82sHhB6TzCW5z5aXoamiKeGy",
		"1dyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9",
		"1PYELM7jXHy5HhatbXGXfRpGrgMMxmpobu",
		"17abzUBJr7cnqfnxnmznn8W38s9f9EoXiq",
		"1DMGtVnRrgZaji7C9noZS3a1QtoaAN2uRG",
		"1CYG7y3fukVLdobqgUtbknwWKUZ5p1HVmV",
	}

	response, e := New().GetUnspentAdvanced(addresses, map[string]string{"limit": "1000"})
	if e != nil {
		t.Fatal(e)
	}

	for _, v := range response.UnspentOutputs {
		if len(v.TxHash) != 64 {
			t.Fatal("Wrong length value on field 'TxHash'")
		}

		if v.TxIndex < 1 {
			t.Fatal("Wrong value on field 'TxIndex'")
		}

		if v.Script == "" {
			t.Fatal("Wrong count items on field 'Script'")
		}
	}
}

// Copyright 2017-2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import (
	"encoding/json"
	"testing"
)

func TestGetChart(t *testing.T) {
	chart, e := newClient().GetChart("market-price")
	if e != nil {
		t.Fatal(e)
	}

	for k, v := range chart.Values {
		if v.X == 0 {
			t.Fatal("Zero time, it is a mistake")
		}
		if v.X > 1500000000 && v.Y == 0 {
			t.Fatal("Error parsing value for Y")
		}
		t.Log(k, v.X, v.Y)
	}
}

func TestGetChartPools(t *testing.T) {
	pools, e := newClient().GetPools()
	if e != nil {
		t.Fatal(e)
	}

	for k, v := range pools {
		t.Logf("%s: %d", k, v)
	}
}

func TestGetStats(t *testing.T) {
	stats, e := newClient().GetStats()
	if e != nil {
		t.Fatal(e)
	}

	bytes, e2 := json.MarshalIndent(stats, "", "\t")
	if e2 != nil {
		t.Fatal(e2)
	}
	t.Log(string(bytes))
}

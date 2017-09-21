// Copyright 2017 Vasilyuk Vasiliy. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import "testing"

func TestGetChart(t *testing.T) {
	c := New()
	chart, e := c.GetChart("market-price")
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
	c := New()
	pools, e := c.GetPools()
	if e != nil {
		t.Fatal(e)
	}

	for k, v := range *pools {
		t.Logf("%s: %d", k, v)
	}
}

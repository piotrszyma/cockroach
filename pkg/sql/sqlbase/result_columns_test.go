// Copyright 2017 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package sqlbase

import (
	"fmt"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/sql/sem/types"
)

func TestResultColumnsTypesEqual(t *testing.T) {
	tests := []struct {
		r, o  ResultColumns
		equal bool
	}{
		{
			r:     ResultColumns{{Typ: types.TypeInt}},
			o:     ResultColumns{{Typ: types.TypeInt}},
			equal: true,
		},
		{
			r:     ResultColumns{{Typ: types.TypeInt}},
			o:     ResultColumns{{Typ: types.TypeString}},
			equal: false,
		},
		{
			r:     ResultColumns{{Typ: types.TypeNull}},
			o:     ResultColumns{{Typ: types.TypeInt}},
			equal: false,
		},
		{
			r:     ResultColumns{{Typ: types.TypeInt}},
			o:     ResultColumns{{Typ: types.TypeNull}},
			equal: true,
		},
		{
			r:     ResultColumns{{Typ: types.TypeNull}},
			o:     ResultColumns{{Typ: types.TypeNull}},
			equal: true,
		},
		{
			r:     ResultColumns{{Typ: types.TypeInt}, {Typ: types.TypeInt}},
			o:     ResultColumns{{Typ: types.TypeInt}},
			equal: false,
		},
		{
			r:     ResultColumns{},
			o:     ResultColumns{{Typ: types.TypeNull}},
			equal: false,
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v-%v", tc.r, tc.o), func(t *testing.T) {
			eq := tc.r.TypesEqual(tc.o)
			if eq != tc.equal {
				t.Fatalf("expected %v, got %v", tc.equal, eq)
			}
		})
	}
}

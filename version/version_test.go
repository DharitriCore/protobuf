// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package version

import "testing"

func TestLess(t *testing.T) {
	tests := []struct {
		name string
		this string
		that string
		want bool
	}{
		{
			name: "equal three component versions",
			this: "3.0.0",
			that: "3.0.0",
			want: true,
		},
		{
			name: "modern protoc two component version",
			this: "3.0.0",
			that: "25.3",
			want: true,
		},
		{
			name: "missing patch is zero",
			this: "25.3.0",
			that: "25.3",
			want: true,
		},
		{
			name: "multi digit minor compares numerically",
			this: "3.10.0",
			that: "3.9.9",
			want: false,
		},
		{
			name: "release suffix is ignored",
			this: "25.3.0",
			that: "25.3.0-rc1",
			want: true,
		},
		{
			name: "higher patch passes",
			this: "3.0.0",
			that: "3.0.1",
			want: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := less(test.this, test.that); got != test.want {
				t.Fatalf("less(%q, %q) = %t, want %t", test.this, test.that, got, test.want)
			}
		})
	}
}

func TestParseVersionRejectsInvalidVersions(t *testing.T) {
	tests := []string{
		"25",
		"25.",
		"25.beta",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			if _, err := parseVersion(test); err == nil {
				t.Fatalf("parseVersion(%q) succeeded, want error", test)
			}
		})
	}
}

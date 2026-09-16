// SPDX-FileCopyrightText: INDUSTRIA DE DISEÑO TEXTIL S.A. (INDITEX S.A.)
// SPDX-License-Identifier: Apache-2.0

package main

import "testing"

func TestGreeting(t *testing.T) {
	if greeting() == "" {
		t.Fatal("greeting must not be empty")
	}
}

func TestFarewell(t *testing.T) {
	if farewell() == "" {
		t.Fatal("farewell must not be empty")
	}
}

func FuzzMessage(f *testing.F) {
	f.Add("Hello")
	f.Add("Goodbye")
	f.Add("")

	f.Fuzz(func(t *testing.T, prefix string) {
		got := message(prefix)
		if len(got) != len(prefix)+len(" from go-ci-testing") {
			t.Fatalf("message(%q) returned unexpected length: %q", prefix, got)
		}
	})
}

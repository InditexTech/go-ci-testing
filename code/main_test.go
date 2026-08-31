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

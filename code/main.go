// SPDX-FileCopyrightText: INDUSTRIA DE DISEÑO TEXTIL S.A. (INDITEX S.A.)
// SPDX-License-Identifier: Apache-2.0

package main

import "fmt"

func greeting() string {
	return message("Hello")
}

func farewell() string {
	return message("Goodbye")
}

func message(prefix string) string {
	return prefix + " from go-ci-testing"
}

func main() {
	fmt.Println(greeting())
	fmt.Println(farewell())
}

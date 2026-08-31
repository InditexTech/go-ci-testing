// SPDX-FileCopyrightText: INDUSTRIA DE DISEÑO TEXTIL S.A. (INDITEX S.A.)
// SPDX-License-Identifier: Apache-2.0

package main

import "fmt"

func greeting() string {
	return "Hello from go-ci-testing"
}

func farewell() string {
	return "Goodbye from go-ci-testing"
}

func main() {
	fmt.Println(greeting())
	fmt.Println(farewell())
}

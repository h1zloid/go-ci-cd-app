package main

import "testing"

func TestDummy(t *testing.T) {
    if 2+2 != 4 {
        t.Error("Math is broken")
    }
}


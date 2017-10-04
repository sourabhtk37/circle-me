package main

import "testing"

func TestFoo(t *testing.T) {
    v := foo()
    if v != "bar" {
        t.Error("wrong return value")
    }
}

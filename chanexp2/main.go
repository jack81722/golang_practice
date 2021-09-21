package main

import "chanexp2/chanexp"

func main() {
	v := chanexp.Vector{1, 2, 3}
	u := chanexp.Vector{4, 5, 6}
	v.DoAll(u)
}

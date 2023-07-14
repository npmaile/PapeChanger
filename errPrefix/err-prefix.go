package errprefix

import "math/rand"

var prefixes = []string{
	"My brother in Christ, your computer has failed: ",
	"UwU OwO x3 OopSiE: ",
	"whoopsie dasies: ",
	"H*ck",
	"😭😭oOpSy DoOpSiE, you made a frickey-wickey 😭😭: ",
	"segmentation fault (core dumped)... JUST KIDDING, but really the program failed",
	"Damn bro, I'm sorry but: ",

}

func Get() string {
	index := rand.Int() % len(prefixes)
	return prefixes[index]

}

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
	"心からのお詫び: ",
  "I'm sorry, I relly fucked up here : ",
  "Sucks to be you: ",
  "Sorry not sorry: ",
  "I'll get back to you soon my bad: ",
  "Not even me am know: ",
  "What the hell did you even do: ",
  "God, nate relly screwed this one up: ",
  "Out Of Memory Error: ",
  "Sorry, I think I had a little bit too much to drink: ",
  "I'm feeling a bit sick today, i'm not gonna make it boss: ",
  "Hey, we messed up, we'll go ahead and submit a issue to the github: ",
  "Sorry I'm having a miragraine, can't make it today: ",





}

func Get() string {
	index := rand.Int() % len(prefixes)
	return prefixes[index]

}

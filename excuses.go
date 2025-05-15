package main

import "math/rand"

var excuses = []string{
	"My cat deployed to production by accident.",
	"It’s not a bug, it’s an undocumented feature.",
	"The internet gods were not pleased today.",
	"I thought you said next sprint.",
	"I'm waiting for Mercury to be out of retrograde.",
	"It worked on my machine™.",
	"My pet Ted Lasso refused to come out of the pitch, so I couldn't leave the house.",
	"A rogue squirrel hacked my keyboard and scrambled my code into hieroglyphs.",
	"My time-traveling debugger sent my variables back to 1985; I'm chasing them across decades.",
	"Aliens abducted my desk and returned it without my monitors.",
	"I accidentally set my coffee mug on fire and had to call the fire brigade.",
	"My code got jealous and started throwing exceptions at me relentlessly.",
	"The ghost in my machine locked me out with a digital padlock.",
	"My goldfish swallowed my USB drive, and I’m waiting for it to… pass the data.",
	"A random wizard cast a latency spell on my network; I'm negotiating an antidote.",
	"My Wi-Fi router staged a coup and declared itself sovereign; it demands tribute.",
	"My rubber duck went on strike until I fix the bug.",
}

func getRandomExcuse() string {
	return excuses[rand.Intn(len(excuses))]
}

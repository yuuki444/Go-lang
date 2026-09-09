package main

import (
	"sync"
)

var (
	movies = []Movie{
		{
			ID:       1,
			Title:    "Interstellar",
			Year:     2014,
			Rating:   8.7,
			Director: "Christopher Nolan",
		},
		{
			ID:       2,
			Title:    "Dune",
			Year:     2021,
			Rating:   8.0,
			Director: "Denis Villeneuve",
		},
		{
			ID:       3,
			Title:    "The Matrix",
			Year:     1999,
			Rating:   8.7,
			Director: "Wachowski",
		},
	}
	mu     sync.Mutex
	nextID = 4
)

package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ShowResult(numbers string) string {
	b := a.digitChecker(numbers)
	var s string
	if b == true {
		s = "OK"
	} else {
		s = "NO"
	}
	return s
}

func (a *App) digitChecker(numbers string) bool {

	sarr := strings.Split(numbers, "")
	var array [12]int

	for i, s := range sarr {
		element, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("Error occured while convert string to interger")
		}
		array[i] = element
	}

	var sum int
	forcheck := array[0:10]
	checker := []int{5, 4, 3, 2, 7, 6, 5, 4, 3, 2}
	for i, s := range forcheck {
		sum += s * checker[i]
	}
	m := sum % 11
	c := 11 - m
	b := false
	if array[10] == c {
		b = true
	}
	return b
}

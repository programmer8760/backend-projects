package main

import (
	"context"
	"fmt"
	"math"
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

func (a *App) ConvertLength(length float64, from int, to int) string {
	scales := [8]float64{0.001, 0.01, 1, 1000, 0.0254, 0.30478, 0.91407, 0.0006215}
	unitsNames := [8]string{"mm", "cm", "m", "km", "in", "ft", "yd", "mi"}
	return fmt.Sprintf("%v%s = %v%s", length, unitsNames[from], math.Round(length*scales[from]/scales[to]*10000)/10000, unitsNames[to])
}

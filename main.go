// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Example - Demonstrates receiving window events.
package main

import (
	"azul3d.org/gfx.v1"
	"azul3d.org/gfx/window.v2"
	"azul3d.org/keyboard.v1"
	_ "azul3d.org/mouse.v1"
	"fmt"
	"image"
	"movement/entity"
	"movement/listener"
	"reflect"
	_ "time"
)

var RenderEntity map[string]entity.Entity = make(map[string]entity.Entity)
var keylistener *listener.KeyboardListener = listener.NewKeyListener()
var statelistener *listener.PositionListener = listener.NewPositionListener()

// gfxLoop is responsible for drawing things to the window.
func gfxLoop(w window.Window, r gfx.Renderer) {
	// You can handle window events in a seperate goroutine!
	go func() {
		// Create our events channel with sufficient buffer size.
		events := make(chan window.Event, 256)

		// Notify our channel anytime any event occurs.
		w.Notify(events, window.AllEvents)

		// Wait for events.
		for event := range events {
			switch ev := event.(type) {
			//case keyboard.TypedEvent:
			case keyboard.StateEvent:
				keylistener.ProcessKey(ev)
				// Use reflection to print the type of event:
				fmt.Println("Event type:", reflect.TypeOf(event))

				// Print the event:
				fmt.Println(event)
			default:
			}
		}
	}()
	for {
		// Clear the entire area (empty rectangle means "the whole area").
		r.Clear(image.Rect(0, 0, 0, 0), gfx.Color{1, 1, 1, 1})
		/*
			// The keyboard is monitored for you, simply check if a key is down:
			if w.Keyboard().Down(keyboard.Space) {
				// Clear a red rectangle.
				r.Clear(image.Rect(0, 0, 100, 100), gfx.Color{1, 0, 0, 1})
			}

			// And the same thing with the mouse, check if a mouse button is down:
			if w.Mouse().Down(mouse.Left) {
				// Clear a blue rectangle.
				r.Clear(image.Rect(100, 100, 200, 200), gfx.Color{0, 0, 1, 1})
			}
		*/
		statelistener.Process()
		drawAll(r)
		// Render the whole frame.
		r.Render()
	}
}

//allow a shape to be drawn on the screen
func RegisterForRender(ent entity.Entity) {
	//only adds to the list if
	if _, found := RenderEntity[ent.Id()]; !found {
		RenderEntity[ent.Id()] = ent
	}
}

//delete the shape from the render list
func UnRegisterForRender(ent entity.Entity) {
	delete(RenderEntity, ent.Id())
}

func drawAll(r gfx.Renderer) {
	r.Clear(image.Rect(0, 0, 0, 0), gfx.Color{1, 1, 1, 1})
	for _, v := range RenderEntity {
		//in this case v.Draw will return both a Rectangle, and a Color
		r.Clear(v.DrawInfo())
	}
}
func main() {
	//Registering the Player person
	x := entity.NewPlayer("hello", image.Rect(100, 100, 150, 150), gfx.Color{0, 0, 1, 1})
	RegisterForRender(x)
	keylistener.Register(x)
	statelistener.RegisterForState(x)

	//adding flooring
	flooring := entity.Wall{entity.NewSimple("flooring", image.Rect(0, 400, 800, 425), gfx.Color{0, 0, 0, 1})}
	RegisterForRender(flooring)
	statelistener.RegisterForState(flooring)

	roof := entity.Wall{entity.NewSimple("roof", image.Rect(0, 0, 800, 25), gfx.Color{0, 0, 0, 1})}
	RegisterForRender(roof)
	statelistener.RegisterForState(roof)

	wall := entity.Wall{entity.NewSimple("wall", image.Rect(500, 0, 525, 400), gfx.Color{0, 0, 0, 1})}
	RegisterForRender(wall)
	statelistener.RegisterForState(wall)

	enemy := entity.NewEnemyPatrol(entity.NewSimple("enemy", image.Rect(250, 250, 300, 300), gfx.Color{1, 0, 0, 1}), image.Pt(100, 275), image.Pt(500, 400))
	RegisterForRender(enemy)
	statelistener.RegisterForState(enemy)
	window.Run(gfxLoop, nil)
}

package pattern

import "fmt"

/*
Implementation of "state" pattern - behavioral pattern (https://en.wikipedia.org/wiki/State_pattern)

Applicability of state pattern:
1. Object Behavior Depends on its State
2. State-Specific Behavior
There are multiple states, and each state encapsulates specific behavior related to that state

Pros:
1. Open/Closed Principle
Adding new states or changing existing ones doesn't require modifying the context class.
This aligns with the Open/Closed Principle
2. Simplifies Context Code
The context (object whose behavior changes) doesn't need to know the details of state transitions.
It interacts with the current state, simplifying its code

Cons:
1. Complexity in Small Systems
In small systems, the overhead of creating multiple state classes might outweigh the benefits of using the pattern.

Real examples of state pattern:
1. Document Editing in a Text Editor
2. Traffic Light System

*/

// State - common interface for various states
type State interface {
	Play()
	Next()
	Previous()
}

// PlayState - concrete state
type PlayState struct{}

func (ps *PlayState) Play() {
	fmt.Println("Playing the current track")
}

func (ps *PlayState) Next() {
	fmt.Println("Playing the next track")
}

func (ps *PlayState) Previous() {
	fmt.Println("Playing the previous track")
}

// NextState - concrete state next track
type NextState struct{}

func (ns *NextState) Play() {
	fmt.Println("Playing the next track")
}

func (ns *NextState) Next() {
	fmt.Println("Playing the next track")
}

func (ns *NextState) Previous() {
	fmt.Println("Playing the previous track")
}

// PreviousState - concrete state previous track
type PreviousState struct{}

func (ps *PreviousState) Play() {
	fmt.Println("Playing the previous track")
}

func (ps *PreviousState) Next() {
	fmt.Println("Playing the next track")
}

func (ps *PreviousState) Previous() {
	fmt.Println("Playing the previous track")
}

// context
type AudioPlayer struct {
	state State
}

func NewAudioPlayer(initialState State) *AudioPlayer {
	return &AudioPlayer{state: initialState}
}
func (ap *AudioPlayer) SetState(newState State) {
	ap.state = newState
}

func (ap *AudioPlayer) Play() {
	ap.state.Play()
}

func (ap *AudioPlayer) Next() {
	ap.state.Next()
}

func (ap *AudioPlayer) Previous() {
	ap.state.Previous()
}

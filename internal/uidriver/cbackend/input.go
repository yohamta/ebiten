// Copyright 2021 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build ebitencbackend
// +build ebitencbackend

package cbackend

import (
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2/internal/cbackend"
	"github.com/hajimehoshi/ebiten/v2/internal/driver"
)

type Input struct {
	touches []cbackend.Touch

	m sync.Mutex
}

func (i *Input) update(context driver.UIContext) {
	i.m.Lock()
	defer i.m.Unlock()

	i.touches = i.touches[:0]
	i.touches = cbackend.AppendTouches(i.touches)

	for idx, t := range i.touches {
		x, y := context.AdjustPosition(float64(t.X), float64(t.Y), 1)
		i.touches[idx].X = int(x)
		i.touches[idx].Y = int(y)
	}
}

func (i *Input) AppendInputChars(runes []rune) []rune {
	return nil
}

func (i *Input) AppendGamepadIDs(gamepadIDs []driver.GamepadID) []driver.GamepadID {
	return nil
}

func (i *Input) AppendTouchIDs(touchIDs []driver.TouchID) []driver.TouchID {
	i.m.Lock()
	defer i.m.Unlock()

	for _, t := range i.touches {
		touchIDs = append(touchIDs, t.ID)
	}
	return touchIDs
}

func (i *Input) CursorPosition() (x, y int) {
	return 0, 0
}

func (i *Input) GamepadSDLID(id driver.GamepadID) string {
	return ""
}

func (i *Input) GamepadName(id driver.GamepadID) string {
	return ""
}

func (i *Input) GamepadAxisValue(id driver.GamepadID, axis int) float64 {
	return 0
}

func (i *Input) GamepadAxisNum(id driver.GamepadID) int {
	return 0
}

func (i *Input) GamepadButtonNum(id driver.GamepadID) int {
	return 0
}

func (i *Input) IsGamepadButtonPressed(id driver.GamepadID, button driver.GamepadButton) bool {
	return false
}

func (i *Input) IsKeyPressed(key driver.Key) bool {
	return false
}

func (i *Input) IsMouseButtonPressed(button driver.MouseButton) bool {
	return false
}

func (i *Input) IsStandardGamepadButtonPressed(id driver.GamepadID, button driver.StandardGamepadButton) bool {
	return false
}

func (i *Input) IsStandardGamepadLayoutAvailable(id driver.GamepadID) bool {
	return true
}

func (i *Input) StandardGamepadAxisValue(id driver.GamepadID, button driver.StandardGamepadAxis) float64 {
	return 0
}

func (i *Input) StandardGamepadButtonValue(id driver.GamepadID, button driver.StandardGamepadButton) float64 {
	return 0
}

func (i *Input) TouchPosition(id driver.TouchID) (x, y int) {
	i.m.Lock()
	defer i.m.Unlock()

	for _, t := range i.touches {
		if t.ID == id {
			return t.X, t.Y
		}
	}
	return 0, 0
}

func (i *Input) VibrateGamepad(id driver.GamepadID, duration time.Duration, strongMagnitude float64, weakMagnitude float64) {
}

func (i *Input) Wheel() (xoff, yoff float64) {
	return 0, 0
}

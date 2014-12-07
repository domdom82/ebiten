package glfw

import (
	glfw "github.com/go-gl/glfw3"
	"github.com/hajimehoshi/ebiten/input"
)

type Keyboard struct {
	pressedKeys map[input.Key]struct{}
}

func NewKeyboard() *Keyboard {
	return &Keyboard{
		pressedKeys: map[input.Key]struct{}{},
	}
}

func (k *Keyboard) IsKeyPressed(key input.Key) bool {
	_, ok := k.pressedKeys[key]
	return ok
}

var glfwKeyCodeToKey = map[glfw.Key]input.Key{
	glfw.KeySpace: input.KeySpace,
	glfw.KeyLeft:  input.KeyLeft,
	glfw.KeyRight: input.KeyRight,
	glfw.KeyUp:    input.KeyUp,
	glfw.KeyDown:  input.KeyDown,
}

func (k *Keyboard) update(window *glfw.Window) {
	for g, u := range glfwKeyCodeToKey {
		if window.GetKey(g) == glfw.Press {
			k.pressedKeys[u] = struct{}{}
		} else {
			delete(k.pressedKeys, u)
		}
	}
}

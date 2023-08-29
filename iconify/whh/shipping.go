package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shipping(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 640h-82q18 31 18 64q0 53-37.5 90.5T768 832t-90.5-37.5T640 704q0-33 18-64H366q18 31 18 64q0 53-37.5 90.5T256 832t-90.5-37.5T128 704q0-33 18-64H64q-26 0-45-18.5T0 576V322q-1-29 19-50l125-125q22-21 53-19h187V64q0-27 19-45.5T448 0h512q26 0 45 19t19 45v512q0 26-18.5 45T960 640zM384 224q0-13-9.5-22.5T352 192H195L64 323v93q0 13 9.5 22.5T96 448h256q13 0 22.5-9.5T384 416V224z"/>`),
		g.Group(children),
	)
}
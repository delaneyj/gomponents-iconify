package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppleKeyboardControl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.78 11.78l-1.42 1.41L12 6.83l-6.36 6.36l-1.42-1.41L12 4l7.78 7.78Z"/>`),
		g.Group(children),
	)
}
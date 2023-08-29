package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Directions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 14.5V12h-4v3H8v-4a1 1 0 0 1 1-1h5V7.5l3.5 3.5m4.21.29l-9-9h-.01a.996.996 0 0 0-1.41 0l-9 9c-.39.39-.39 1.03 0 1.42l9 9c.39.38 1.02.39 1.42 0l9-9c.39-.39.39-1.03 0-1.42Z"/>`),
		g.Group(children),
	)
}
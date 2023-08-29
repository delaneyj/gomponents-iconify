package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BreadSlice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.5 0 10 3.36 10 7.5c0 1.69-.74 3.25-2 4.5v8H4v-8c-1.26-1.25-2-2.81-2-4.5C2 5.36 6.5 2 12 2M8 18h4v-4H8v4Z"/>`),
		g.Group(children),
	)
}
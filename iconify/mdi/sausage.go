package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sausage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 5.3c1.2.5 2 1.8 2 3.2C21 15.4 15.4 21 8.5 21c-1.4 0-2.6-.8-3.2-2L3 20.5v-6L5.3 16c.6-1.2 1.8-2 3.2-2c3 0 5.5-2.5 5.5-5.5c0-1.4.8-2.6 2-3.2L14.5 3h6L19 5.3Z"/>`),
		g.Group(children),
	)
}
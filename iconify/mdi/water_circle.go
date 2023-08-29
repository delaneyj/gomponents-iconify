package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m0 17c-2.76 0-5-2.24-5-5c0-3.33 5-8.96 5-8.96s5 5.63 5 8.96c0 2.76-2.24 5-5 5Z"/>`),
		g.Group(children),
	)
}
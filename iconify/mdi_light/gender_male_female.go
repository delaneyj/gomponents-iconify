package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderMaleFemale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 3v1h3.3L15 8.3C14 7.5 12.8 7 11.5 7C8.5 7 6 9.5 6 12.5c0 2.8 2.2 5.2 5 5.5v2H9v1h2v2h1v-2h2v-1h-2v-2c2.8-.3 5-2.6 5-5.5c0-1.3-.5-2.5-1.3-3.5L20 4.7V8h1V3h-5m-4.5 5c1 0 2 .4 2.8 1c.2.2.5.4.7.7c.6.8 1 1.8 1 2.8c0 2.5-2 4.5-4.5 4.5S7 15 7 12.5S9 8 11.5 8Z"/>`),
		g.Group(children),
	)
}
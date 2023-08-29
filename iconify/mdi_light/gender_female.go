package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderFemale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 5C8.5 5 6 7.5 6 10.5c0 2.8 2.2 5.2 5 5.5v2H9v1h2v2h1v-2h2v-1h-2v-2c2.8-.3 5-2.6 5-5.5c0-3-2.5-5.5-5.5-5.5m0 1C14 6 16 8 16 10.5S14 15 11.5 15S7 13 7 10.5S9 6 11.5 6Z"/>`),
		g.Group(children),
	)
}
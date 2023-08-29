package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalCottonSwab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h12v2H3V3M2 21h14V6H2v15m3-9h2.5V9.5h3V12H13v3h-2.5v2.5h-3V15H5v-3m15-6c-1.7 0-3 1.8-3 4c0 1.8.8 3.2 2 3.8V21h2v-7.2c1.2-.5 2-2 2-3.8c0-2.2-1.3-4-3-4Z"/>`),
		g.Group(children),
	)
}
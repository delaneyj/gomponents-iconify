package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighwayRestArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13.5 13h-3V9h3a.5.5 0 0 0 .41-.787L11.66 5h.84a.5.5 0 0 0 .384-.82l-2.5-3a.515.515 0 0 0-.768 0l-2.5 3A.5.5 0 0 0 7.5 5h.84L6.09 8.213A.5.5 0 0 0 6.5 9h3v4H4v-2h1.5a.5.5 0 0 0 0-1h-4a.5.5 0 0 0 0 1H3v2H1.5a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1Z"/>`),
		g.Group(children),
	)
}
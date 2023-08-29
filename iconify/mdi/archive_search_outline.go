package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveSearchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.04 10c-.4.25-.78.55-1.14.9c-.33.34-.63.71-.87 1.1H8v-1.5c0-.28.22-.5.5-.5h4.54M20 8H2V2h18v6m-2-4H4v2h14V4M5 18V9H3v11h8.82a6.44 6.44 0 0 1-1.32-2H5m18.39 3L22 22.39l-3.12-3.07c-.69.43-1.51.68-2.38.68c-2.5 0-4.5-2-4.5-4.5s2-4.5 4.5-4.5s4.5 2 4.5 4.5c0 .88-.25 1.71-.69 2.4l3.08 3.1M19 15.5a2.5 2.5 0 0 0-5 0a2.5 2.5 0 0 0 5 0Z"/>`),
		g.Group(children),
	)
}
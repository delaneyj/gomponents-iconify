package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarExport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 22l4-4h-3v-6h-2v6H8l4 4m7-19h-1V1h-2v2H8V1H6v2H5c-1.1 0-2 .9-2 2v14a2 2 0 0 0 2 2h3l-2-2H5V8h14v11h-1l-2 2h3c1.11 0 2-.89 2-2V5a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}
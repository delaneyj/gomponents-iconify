package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 13c.34 0 .67.04 1 .09V8H4v13h9.35c-.22-.63-.35-1.3-.35-2c0-3.31 2.69-6 6-6M9 13v-1.5c0-.28.22-.5.5-.5h5c.28 0 .5.22.5.5V13H9m12-6H3V3h18v4m2 11v2h-8v-2h8Z"/>`),
		g.Group(children),
	)
}
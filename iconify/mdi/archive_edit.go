package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 10.3V8H4v13h7v-1.87l8.39-8.39c.18-.18.39-.32.61-.44M15 13H9v-1.5c0-.28.22-.5.5-.5h5c.28 0 .5.22.5.5V13m6-6H3V3h18v4m1.85 7.19l-.98.98l-2.04-2.04l.98-.98c.19-.2.52-.2.72 0l1.32 1.32c.2.2.2.53 0 .72m-3.72-.36l2.04 2.04L15.04 22H13v-2.04l6.13-6.13Z"/>`),
		g.Group(children),
	)
}
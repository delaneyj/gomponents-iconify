package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RenameOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15 16l-4 4h10v-4h-6m-2.94-8.81L3 16.25V20h3.75l9.06-9.06l-3.75-3.75M5.92 18H5v-.92L12.06 10l.94.94L5.92 18m12.79-9.96c.39-.39.39-1.04 0-1.41l-2.34-2.34a1.001 1.001 0 0 0-1.41 0l-1.83 1.83l3.75 3.75l1.83-1.83Z"/>`),
		g.Group(children),
	)
}
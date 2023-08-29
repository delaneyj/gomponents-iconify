package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1.198l10 8.334V22H2V9.532l10-8.334ZM10 20h4v-5h-4v5Zm6 0h4v-9.532l-8-6.666l-8 6.666V20h4v-7h8v7Z"/>`),
		g.Group(children),
	)
}
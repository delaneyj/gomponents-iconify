package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RangeHoodOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18v-4.2q0-.4.163-.763T2.6 12.4L7 8V3h10v5l4.425 4.425q.275.275.425.638t.15.762V18q0 .825-.588 1.413T20 20H4Zm1.8-8h12.4L15 8.8V5H9v3.8L5.8 12ZM4 18h16v-4H4v4Zm6-1.3v-1.5h4v1.5h-4Z"/>`),
		g.Group(children),
	)
}
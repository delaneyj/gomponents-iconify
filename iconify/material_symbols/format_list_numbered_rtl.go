package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatListNumberedRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 22v-1.5h2.5v-.75H18v-1.5h1.5v-.75H17V16h3q.425 0 .713.288T21 17v1q0 .425-.288.713T20 19q.425 0 .713.288T21 20v1q0 .425-.288.713T20 22h-3Zm0-7v-2.75q0-.425.288-.713T18 11.25h1.5v-.75H17V9h3q.425 0 .713.288T21 10v1.75q0 .425-.288.713T20 12.75h-1.5v.75H21V15h-4Zm1.5-7V3.5H17V2h3v6h-1.5ZM3 19v-2h12v2H3Zm0-6v-2h12v2H3Zm0-6V5h12v2H3Z"/>`),
		g.Group(children),
	)
}
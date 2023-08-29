package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OfflineShareOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 23V5h2v16h10v2H4Zm4-4V1h12v18H8Zm2-3v1h8v-1h-8Zm0-2h8V6h-8v8Zm1-2V8.75h3.15l-.7-.7L14.5 7L17 9.5L14.5 12l-1.05-1.05l.7-.7H12.5V12H11Zm-1-8h8V3h-8v1Zm0 0V3v1Zm0 12v1v-1Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomainDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 19.15l-2-2V9h-8.15L10 7.15V5H7.85l-2-2H12v4h10v12.15Zm-1.5 4.15L18.15 21H2V4.8L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM4 19h2v-2H4v2Zm0-4h2v-2H4v2Zm0-4h2V9H4v2Zm4 8h2v-2H8v2Zm0-4h2v-2H8v2Zm4 4h4.15l-2-2H12v2Zm6-6h-2v-2h2v2Z"/>`),
		g.Group(children),
	)
}
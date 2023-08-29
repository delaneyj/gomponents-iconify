package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabPanelOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21v-9h2V8H2V3h9v5h-1v4h4V8h-1V3h9v5h-1v4h2v9H1ZM15 6h5V5h-5v1ZM4 6h5V5H4v1Zm12 6h3V8h-3v4ZM5 12h3V8H5v4Zm-2 7h18v-5H3v5ZM4 6V5v1Zm11 0V5v1ZM3 19v-5v5Z"/>`),
		g.Group(children),
	)
}
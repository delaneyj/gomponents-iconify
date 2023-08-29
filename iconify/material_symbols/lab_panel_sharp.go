package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabPanelSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21v-9h2V8H2V3h9v5h-1v4h4V8h-1V3h9v5h-1v4h2v9H1Zm15-9h3V8h-3v4ZM5 12h3V8H5v4Z"/>`),
		g.Group(children),
	)
}
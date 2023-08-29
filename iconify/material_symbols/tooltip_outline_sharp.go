package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TooltipOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 22l-2.65-4H2V2h20v16h-7.35L12 22Zm-8-6h16V4H4v12Zm0 0V4v12Z"/>`),
		g.Group(children),
	)
}
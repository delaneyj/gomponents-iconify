package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.3 15.3l-2.6-2.6q-.15-.15-.225-.325T9.4 12q0-.2.075-.375T9.7 11.3l2.6-2.6q.475-.475 1.088-.212t.612.937v5.15q0 .675-.613.938T12.3 15.3Z"/>`),
		g.Group(children),
	)
}
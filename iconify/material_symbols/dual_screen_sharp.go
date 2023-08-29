package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DualScreenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18.95V2l10 4.025V22.95l-10-4ZM16 19V4.65L9.45 2H20v17h-4Z"/>`),
		g.Group(children),
	)
}
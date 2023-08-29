package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignJustifyFlexEnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 22V2h2v20h-2Zm-6-5V7h3v10h-3Zm-6 0V7h3v10H8Z"/>`),
		g.Group(children),
	)
}
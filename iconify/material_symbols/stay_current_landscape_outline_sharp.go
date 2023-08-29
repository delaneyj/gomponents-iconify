package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StayCurrentLandscapeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 19V5h22v14H1Zm3-2V7H3v10h1Zm2 0h12V7H6v10Zm14 0h1V7h-1v10ZM4 7H3h1Zm16 0h1h-1Z"/>`),
		g.Group(children),
	)
}
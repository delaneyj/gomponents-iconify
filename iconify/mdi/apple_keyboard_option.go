package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppleKeyboardOption(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h6.11l7.04 14H21v2h-6.12L7.84 6H3V4m11 0h7v2h-7V4Z"/>`),
		g.Group(children),
	)
}
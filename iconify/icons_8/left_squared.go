package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftSquared(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5H5zm2 2h18v18H7V7zm10.78 2.28l-6 6l-.686.72l.687.72l6 6l1.44-1.44L13.937 16l5.28-5.28l-1.437-1.44z"/>`),
		g.Group(children),
	)
}
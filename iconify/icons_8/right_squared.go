package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightSquared(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5H5zm2 2h18v18H7V7zm7.22 2.28l-1.44 1.44L18.064 16l-5.282 5.28l1.44 1.44l6-6l.686-.72l-.687-.72l-6-6z"/>`),
		g.Group(children),
	)
}
package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 13H4a1 1 0 0 1 0-2h16a1 1 0 0 1 0 2Z"/>`),
		g.Group(children),
	)
}
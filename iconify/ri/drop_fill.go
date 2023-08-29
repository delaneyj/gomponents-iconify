package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.636 6.633L12 .269l6.364 6.364a9 9 0 1 1-12.728 0Z"/>`),
		g.Group(children),
	)
}
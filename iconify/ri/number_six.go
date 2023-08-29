package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.886 2l-4.438 7.686A6.5 6.5 0 1 1 6.4 12.7L12.576 2h2.31ZM12 11.5a4.5 4.5 0 1 0 0 9a4.5 4.5 0 0 0 0-9Z"/>`),
		g.Group(children),
	)
}
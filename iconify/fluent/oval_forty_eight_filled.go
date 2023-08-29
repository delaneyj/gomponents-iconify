package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OvalFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 24c0-8.284 6.716-15 15-15h10c8.284 0 15 6.716 15 15c0 8.284-6.716 15-15 15H19c-8.284 0-15-6.716-15-15Z"/>`),
		g.Group(children),
	)
}
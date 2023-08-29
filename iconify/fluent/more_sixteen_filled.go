package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M4 6.75a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5zm4 0a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5zm4 0a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}
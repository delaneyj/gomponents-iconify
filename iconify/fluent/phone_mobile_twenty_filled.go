package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneMobileTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M13.5 2A1.5 1.5 0 0 1 15 3.5v13a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 5 16.5v-13A1.5 1.5 0 0 1 6.5 2zM11 14H9a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}
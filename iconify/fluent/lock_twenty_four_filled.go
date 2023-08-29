package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M12 2a4 4 0 0 1 4 4v2h2.5A1.5 1.5 0 0 1 20 9.5v11a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 4 20.5v-11A1.5 1.5 0 0 1 5.5 8H8V6a4 4 0 0 1 4-4zm0 11.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zM12 4a2 2 0 0 0-2 2v2h4V6a2 2 0 0 0-2-2z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}
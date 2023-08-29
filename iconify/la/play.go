package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 5.156v21.688l1.531-1L25.844 16L10.53 6.156zm2 3.657L22.156 16L11 23.188z"/>`),
		g.Group(children),
	)
}
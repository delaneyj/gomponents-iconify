package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M43.56 4.44a1.5 1.5 0 0 1 0 2.12l-37 37a1.5 1.5 0 0 1-2.12-2.12l37-37a1.5 1.5 0 0 1 2.12 0Z"/>`),
		g.Group(children),
	)
}
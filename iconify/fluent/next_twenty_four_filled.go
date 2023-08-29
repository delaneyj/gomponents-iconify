package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 4.753c0-1.408 1.578-2.24 2.74-1.444l10.498 7.194a1.75 1.75 0 0 1 .01 2.88L5.749 20.685C4.59 21.492 3 20.66 3 19.248V4.753ZM21 3.75a.75.75 0 0 0-1.5 0v16.5a.75.75 0 0 0 1.5 0V3.75Z"/>`),
		g.Group(children),
	)
}
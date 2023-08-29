package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BadgeTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 4.502a3.5 3.5 0 0 0 5.001 3.163L21 18.25A2.75 2.75 0 0 1 18.25 21H5.75A2.75 2.75 0 0 1 3 18.25V5.75A2.75 2.75 0 0 1 5.75 3h10.588A3.486 3.486 0 0 0 16 4.502Zm3.5-2.5a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Z"/>`),
		g.Group(children),
	)
}
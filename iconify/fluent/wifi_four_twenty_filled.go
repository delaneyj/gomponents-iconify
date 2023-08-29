package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiFourTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.921 13.881a1.242 1.242 0 1 1-1.756 1.757a1.242 1.242 0 0 1 1.757-1.757Z"/>`),
		g.Group(children),
	)
}
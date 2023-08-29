package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretRightTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 14.204a1 1 0 0 0 1.628.778l4.723-3.815a1.5 1.5 0 0 0 0-2.334L8.628 5.02A1 1 0 0 0 7 5.797v8.407Z"/>`),
		g.Group(children),
	)
}
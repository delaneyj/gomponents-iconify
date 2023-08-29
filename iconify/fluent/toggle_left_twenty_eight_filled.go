package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleLeftTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 14a6 6 0 0 1 6-6h12a6 6 0 0 1 0 12H8a6 6 0 0 1-6-6Zm7 3a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/>`),
		g.Group(children),
	)
}
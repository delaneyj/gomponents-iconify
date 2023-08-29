package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloverSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 1a3 3 0 1 1 0 6H8.5a.5.5 0 0 1-.5-.5V4a3 3 0 0 1 3-3Zm-7 .02a3 3 0 1 0 0 6h2.5a.5.5 0 0 0 .5-.5v-2.5a3 3 0 0 0-3-3ZM4 14a3 3 0 1 1 0-6h2.5a.5.5 0 0 1 .5.5V11a3 3 0 0 1-3 3Zm7 0a3 3 0 1 0 0-6H8.5a.5.5 0 0 0-.5.5V11a3 3 0 0 0 3 3Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentMultipleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 6a1 1 0 0 0-1-1v6a3 3 0 0 1-3 3H5a1 1 0 0 0 1 1h4.21A3.79 3.79 0 0 0 14 11.21V6ZM7 1v3.5A1.5 1.5 0 0 0 8.5 6h3.502v5.5a1.5 1.5 0 0 1-1.5 1.5H4.5A1.5 1.5 0 0 1 3 11.5v-9A1.5 1.5 0 0 1 4.5 1H7Zm1 .25V4.5a.5.5 0 0 0 .5.5h3.25L8 1.25Z"/>`),
		g.Group(children),
	)
}
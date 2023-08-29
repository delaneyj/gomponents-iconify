package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTreemapTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 3H6a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h1V3Zm1 14h6a3 3 0 0 0 3-3v-1H8v4Zm9-5V6a3 3 0 0 0-3-3H8v9h9Z"/>`),
		g.Group(children),
	)
}
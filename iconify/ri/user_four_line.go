package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserFourLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20h14v2H5v-2Zm7-2a8 8 0 1 1 0-16a8 8 0 0 1 0 16Zm0-2a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/>`),
		g.Group(children),
	)
}
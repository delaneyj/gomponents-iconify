package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelBottomTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 3a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H5ZM3 6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v5H3V6Zm0 6h14v1a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-1Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 4a.5.5 0 1 1 0-1h8a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-1 0V4.707L3.854 16.854a.5.5 0 0 1-.708-.708L15.293 4H8.5Z"/>`),
		g.Group(children),
	)
}
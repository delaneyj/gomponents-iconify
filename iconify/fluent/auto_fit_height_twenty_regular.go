package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoFitHeightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 2a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1H4Zm0 15a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1H4ZM9.146 5.146a.5.5 0 0 1 .708 0l2 2a.5.5 0 0 1-.708.708L10 6.707v6.586l1.146-1.147a.5.5 0 0 1 .708.708l-2 2a.5.5 0 0 1-.708 0l-2-2a.5.5 0 0 1 .708-.708L9 13.293V6.707L7.854 7.854a.5.5 0 1 1-.708-.708l2-2Z"/>`),
		g.Group(children),
	)
}
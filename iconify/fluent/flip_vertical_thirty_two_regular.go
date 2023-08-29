package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVerticalThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M27.54 2.158A1 1 0 0 1 28 3v11a1 1 0 0 1-1 1H3a1 1 0 0 1-.417-1.91l24-11a1 1 0 0 1 .957.068ZM7.582 13H26V4.558L7.582 13ZM28 29a1 1 0 0 1-1.417.91l-24-11A1 1 0 0 1 3 17h24a1 1 0 0 1 1 1v11Z"/>`),
		g.Group(children),
	)
}
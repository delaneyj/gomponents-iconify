package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrameTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 2.5a.5.5 0 0 0-1 0V4H2.5a.5.5 0 0 0 0 1H4v10H2.5a.5.5 0 0 0 0 1H4v1.5a.5.5 0 0 0 1 0V16h10v1.5a.5.5 0 0 0 1 0V16h1.5a.5.5 0 0 0 0-1H16V5h1.5a.5.5 0 0 0 0-1H16V2.5a.5.5 0 0 0-1 0V4H5V2.5ZM15 5v10H5V5h10Z"/>`),
		g.Group(children),
	)
}
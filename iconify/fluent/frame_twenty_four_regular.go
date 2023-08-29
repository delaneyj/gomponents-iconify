package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrameTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 2.75a.75.75 0 0 0-1.5 0V4.5H2.75a.75.75 0 0 0 0 1.5H4.5v12H2.75a.75.75 0 0 0 0 1.5H4.5v1.75a.75.75 0 0 0 1.5 0V19.5h12v1.75a.75.75 0 0 0 1.5 0V19.5h1.75a.75.75 0 0 0 0-1.5H19.5V6h1.75a.75.75 0 0 0 0-1.5H19.5V2.75a.75.75 0 0 0-1.5 0V4.5H6V2.75ZM18 18H6V6h12v12Z"/>`),
		g.Group(children),
	)
}
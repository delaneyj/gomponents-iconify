package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrameTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.25 2a.75.75 0 0 1 .75.75V4.5h12V2.75a.75.75 0 0 1 1.5 0V4.5h1.75a.75.75 0 0 1 0 1.5H19.5v12h1.75a.75.75 0 0 1 0 1.5H19.5v1.75a.75.75 0 0 1-1.5 0V19.5H6v1.75a.75.75 0 0 1-1.5 0V19.5H2.75a.75.75 0 0 1 0-1.5H4.5V6H2.75a.75.75 0 0 1 0-1.5H4.5V2.75A.75.75 0 0 1 5.25 2Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtractFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 23.25c0-.69.56-1.25 1.25-1.25h33.5a1.25 1.25 0 1 1 0 2.5H7.25c-.69 0-1.25-.56-1.25-1.25Z"/>`),
		g.Group(children),
	)
}
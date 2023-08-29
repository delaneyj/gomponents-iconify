package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DustTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm16 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM7 21a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM8.664 6.943A5.252 5.252 0 0 1 19 8.25c0 1.47-.604 2.798-1.576 3.75a5.25 5.25 0 1 1-8.76 5.057a5.25 5.25 0 1 1 0-10.114Z"/>`),
		g.Group(children),
	)
}
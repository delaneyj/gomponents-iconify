package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PreviousTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 4.252c0-1-1.116-1.595-1.947-1.038L6.554 8.921a1.25 1.25 0 0 0-.007 2.071l8.5 5.793A1.25 1.25 0 0 0 17 15.752v-11.5ZM3 3.5a.5.5 0 0 1 1 0v13a.5.5 0 0 1-1 0v-13Z"/>`),
		g.Group(children),
	)
}
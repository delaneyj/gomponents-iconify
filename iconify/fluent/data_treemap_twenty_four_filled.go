package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTreemapTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6.25A3.25 3.25 0 0 1 6.25 3H9v18H6.25A3.25 3.25 0 0 1 3 17.75V6.25ZM10.5 21h7.25A3.25 3.25 0 0 0 21 17.75V15.5H10.5V21ZM21 14V6.25A3.25 3.25 0 0 0 17.75 3H10.5v11H21Z"/>`),
		g.Group(children),
	)
}
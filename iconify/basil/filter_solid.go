package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.523 4.226a58.727 58.727 0 0 0-13.046 0a1.373 1.373 0 0 0-.915 2.229l3.769 4.659A7.5 7.5 0 0 1 10 15.83v3.142a.75.75 0 0 0 .306.605l2.771 2.032a.58.58 0 0 0 .923-.468V15.83a7.5 7.5 0 0 1 1.67-4.717l3.768-4.66a1.373 1.373 0 0 0-.915-2.228Z"/>`),
		g.Group(children),
	)
}
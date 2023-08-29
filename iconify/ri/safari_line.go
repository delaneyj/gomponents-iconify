package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SafariLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.813 6.503l-4.398 6.911l-6.911 4.398A7.973 7.973 0 0 0 11 19.938V18h2v1.938a7.96 7.96 0 0 0 3.906-1.618l-1.37-1.37l1.414-1.414l1.37 1.37A7.96 7.96 0 0 0 19.939 13h-1.938v-2h1.938a7.974 7.974 0 0 0-2.126-4.497Zm-.315-.315a7.973 7.973 0 0 0-4.497-2.126V6h-2V4.062A7.96 7.96 0 0 0 7.095 5.68l1.37 1.37l-1.414 1.414l-1.37-1.37A7.96 7.96 0 0 0 4.063 11H6v2H4.063a7.973 7.973 0 0 0 2.126 4.497l4.398-6.911l6.911-4.398ZM12.001 22c-5.523 0-10-4.477-10-10s4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10Z"/>`),
		g.Group(children),
	)
}
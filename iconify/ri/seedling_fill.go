package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeedlingFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.998 7v2.5a6.5 6.5 0 0 1-6.5 6.5h-2.5v5h-2v-7l.019-1a6.5 6.5 0 0 1 6.481-6h4.5Zm-16-4a7.003 7.003 0 0 1 6.643 4.786A7.48 7.48 0 0 0 10.014 13H8.998a7 7 0 0 1-7-7V3h4Z"/>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 21l6.324-11.69c.54-.974 1.756-4.104-1.499-5.762C10.57 1.891 8.65 4.411 8 5.58"/>`),
		g.Group(children),
	)
}
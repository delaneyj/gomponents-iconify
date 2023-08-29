package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingRetailMoreTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 3.293A1 1 0 0 1 4.707 3h10.586a1 1 0 0 1 .707.293l2.642 2.641c.762.763.222 2.066-.856 2.066H2.214C1.136 8 .597 6.697 1.36 5.934L4 3.293ZM17 9H3v6a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V9ZM7 13a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm3 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm3 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`),
		g.Group(children),
	)
}
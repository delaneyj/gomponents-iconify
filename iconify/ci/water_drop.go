package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterDrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 13.385a4.676 4.676 0 0 1-1.318 3.263A4.473 4.473 0 0 1 13 17.736m6-4.044C19 7.115 12 2 12 2S5 7.115 5 13.692c0 1.938.737 3.797 2.05 5.168C8.363 20.23 10.144 21 12 21c1.857 0 3.637-.77 4.95-2.14c1.313-1.371 2.05-3.23 2.05-5.168Z"/>`),
		g.Group(children),
	)
}
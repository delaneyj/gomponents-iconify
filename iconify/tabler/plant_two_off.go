package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlantTwoOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 9c0 5.523 4.477 10 10 10a9.953 9.953 0 0 0 5.418-1.593m2.137-1.855A9.961 9.961 0 0 0 22 9"/><path d="M12 19c0-1.988.58-3.84 1.58-5.397m1.878-2.167A9.961 9.961 0 0 1 22 9M2 9a10 10 0 0 1 10 10m0-15a9.7 9.7 0 0 1 3 7.013"/><path d="M9.01 11.5a9.696 9.696 0 0 1 .163-2.318m1.082-2.942A9.696 9.696 0 0 1 12 4M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
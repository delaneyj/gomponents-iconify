package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CpuOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 5h9a1 1 0 0 1 1 1v9m-.292 3.706A1 1 0 0 1 18 19H6a1 1 0 0 1-1-1V6c0-.272.108-.518.284-.698"/><path d="M13 9h2v2m0 4H9V9m-6 1h2m-2 4h2m5-11v2m4-2v2m7 5h-2m2 4h-2m-5 7v-2m-4 2v-2M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
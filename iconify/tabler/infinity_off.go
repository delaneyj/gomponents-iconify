package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfinityOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.165 8.174A4 4 0 0 0 2.999 12a4 4 0 0 0 6.829 2.828A10 10 0 0 0 12 12m1.677-2.347a10 10 0 0 1 .495-.481a4 4 0 1 1 5.129 6.1m-3.521.537a4 4 0 0 1-1.608-.981A10 10 0 0 1 12 12M3 3l18 18"/>`),
		g.Group(children),
	)
}
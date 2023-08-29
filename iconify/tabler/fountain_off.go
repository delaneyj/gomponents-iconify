package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FountainOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 16v-5a2 2 0 1 0-4 0m10 5v-1m0-4a2 2 0 1 1 4 0m-7 5v-4m0-4V6a3 3 0 0 1 6 0"/><path d="M7.451 3.43A3 3 0 0 1 12 6m8 10h1v1m-.871 3.114A2.99 2.99 0 0 1 18 21H6a3 3 0 0 1-3-3v-2h13M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
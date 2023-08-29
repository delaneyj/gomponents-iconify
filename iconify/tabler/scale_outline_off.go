package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScaleOutlineOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 3h10a4 4 0 0 1 4 4v10m-1.173 2.83A3.987 3.987 0 0 1 17 21H7a4 4 0 0 1-4-4V7c0-1.104.447-2.103 1.17-2.827"/><path d="M11.062 7.062A7.002 7.002 0 0 1 17 9.095A142.85 142.85 0 0 0 15 11m-3.723.288a3 3 0 0 0-1.315.71L7.006 9.095a6.977 6.977 0 0 1 1.142-.942M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
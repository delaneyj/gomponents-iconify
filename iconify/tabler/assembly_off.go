package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssemblyOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8.703 4.685L11.029 3.3a2.056 2.056 0 0 1 2 0l6 3.573H19a2 2 0 0 1 1 1.747v6.536c0 .248-.046.49-.132.715m-2.156 1.837l-4.741 3.029a2 2 0 0 1-1.942 0l-6-3.833A2 2 0 0 1 4 15.157V8.62a2 2 0 0 1 1.029-1.748l1.157-.689"/><path d="M11.593 7.591c.295-.133.637-.12.921.04l3 1.79H15.5c.312.181.503.516.5.877V12m-1.152 2.86l-2.363 1.514a1 1 0 0 1-.97 0l-3-1.922A1 1 0 0 1 8 13.576v-3.278c0-.364.197-.7.514-.877l.568-.339M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
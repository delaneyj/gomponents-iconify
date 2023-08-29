package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPieAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.25 5.788a7.251 7.251 0 1 0 7.962 7.962H12a.75.75 0 0 1-.75-.75V5.788ZM3.25 13A8.75 8.75 0 0 1 12 4.25a.75.75 0 0 1 .75.75v7.25H20a.75.75 0 0 1 .75.75a8.75 8.75 0 1 1-17.5 0Z"/><path d="M15.5 4.674V9.5h4.826A6.512 6.512 0 0 0 15.5 4.674Zm-.502-1.612c3.62.45 6.49 3.32 6.94 6.94c.069.548-.386.998-.938.998h-6.5a.5.5 0 0 1-.5-.5V4c0-.552.45-1.007.998-.938Z"/></g>`),
		g.Group(children),
	)
}
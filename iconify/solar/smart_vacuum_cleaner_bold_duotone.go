package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartVacuumCleanerBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m16.441 19.427l-1 1.118l.559.5a3.75 3.75 0 0 0 5.295-5.295l-.5-.56l-1.118 1.001l-3.236 3.236Zm-8.882 0l1 1.118l-.559.5a3.75 3.75 0 0 1-5.295-5.295l.5-.56l1.118 1.001l3.236 3.236Z" opacity=".5"/><path d="M12 6.5a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5Z"/><path fill-rule="evenodd" d="M22 11.75c0 5.523-4.477 10-10 10s-10-4.477-10-10s4.477-10 10-10s10 4.477 10 10Zm-13.75-3a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0Zm4.5 7a.75.75 0 0 0-1.5 0v2a.75.75 0 0 0 1.5 0v-2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTubeMinimalisticBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3.187 15.049a4.085 4.085 0 0 0 0 5.758a4.042 4.042 0 0 0 5.734 0l3.746-3.762l-1.772-.736a2.356 2.356 0 0 1-1.408-1.906a2.352 2.352 0 0 0-2.074-2.082h-1.51l-2.716 2.728Z"/><path fill-rule="evenodd" d="M13.363 2.233a.8.8 0 0 1 1.13.003l7.274 7.305a.8.8 0 0 1-1.134 1.129L13.36 3.364a.8.8 0 0 1 .003-1.13Z" clip-rule="evenodd"/><path d="M14.09 4.098L3.187 15.048a4.085 4.085 0 0 0 0 5.76a4.042 4.042 0 0 0 5.734 0L19.824 9.856l-5.734-5.76Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}
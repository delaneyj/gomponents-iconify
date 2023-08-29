package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.91 4.949A9.968 9.968 0 0 0 2 12c0 5.523 4.477 10 10 10a9.968 9.968 0 0 0 7.05-2.909m1.728-2.298A9.955 9.955 0 0 0 22 12c0-5.523-4.477-10-10-10c-1.74 0-3.376.444-4.8 1.225"/><path d="M7 15v-4.5a1.5 1.5 0 0 1 2.138-1.358m.716.711c.094.196.146.415.146.647V15m-3-2h3m4 1v1h1m2-2v-2a2 2 0 0 0-2-2h-1v1M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
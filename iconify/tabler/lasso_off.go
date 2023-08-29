package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LassoOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.028 13.252C3.371 12.28 3 11.174 3 10c0-1.804.878-3.449 2.319-4.69m2.49-1.506A11.066 11.066 0 0 1 12 3c4.97 0 9 3.134 9 7c0 1.799-.873 3.44-2.307 4.68m-2.503 1.517A11.066 11.066 0 0 1 12 17c-1.913 0-3.686-.464-5.144-1.255"/><path d="M3 15a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="M5 17c0 1.42.316 2.805 1 4M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceDeleteThreeRemoveCircleGarbageTrashDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="m2.4 11.6l9.2-9.2m0 9.2L2.4 2.4"/></g>`),
		g.Group(children),
	)
}
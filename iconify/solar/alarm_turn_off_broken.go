package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmTurnOffBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M14.121 15.122L12 13m0 0l-2.121-2.12M12 13l2.121-2.121M12 13l-2.121 2.121"/><path stroke-linejoin="round" d="m3.5 4.5l4-2.5m13 2.5l-4-2.5"/><path d="M7.5 5.204A9 9 0 1 1 4.204 8.5"/></g>`),
		g.Group(children),
	)
}
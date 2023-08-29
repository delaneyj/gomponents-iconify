package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmTurnOffLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="13" r="9"/><path stroke-linecap="round" d="M14.121 15.122L12.002 13m0 0l-2.124-2.12M12 13l2.122-2.121M12 13l-2.12 2.121"/><path stroke-linecap="round" stroke-linejoin="round" d="m3.5 4.5l4-2.5m13 2.5l-4-2.5"/></g>`),
		g.Group(children),
	)
}
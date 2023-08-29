package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicationTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="24" height="10" x="9" y="5" stroke-linecap="round" stroke-linejoin="round" rx="4"/><path stroke-linecap="round" stroke-linejoin="round" d="M21 5v8m-6-8v8m12-8v8m0 27H8a2 2 0 0 1-2-2V17a2 2 0 0 1 2-2h26a2 2 0 0 1 2 2v9m-2 6v4h4"/><circle cx="35" cy="35" r="9"/><path stroke-linecap="round" stroke-linejoin="round" d="M13 5h16M13 15h16"/></g>`),
		g.Group(children),
	)
}
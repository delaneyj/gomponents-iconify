package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChecklistLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M2 5.5L3.214 7L7.5 3"/><path stroke-linejoin="round" d="M2 12.5L3.214 14L7.5 10" opacity=".5"/><path stroke-linejoin="round" d="M2 19.5L3.214 21L7.5 17"/><path d="M22 19H12"/><path d="M22 12H12" opacity=".5"/><path d="M22 5H12"/></g>`),
		g.Group(children),
	)
}
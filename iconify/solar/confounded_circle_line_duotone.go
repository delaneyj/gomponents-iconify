package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConfoundedCircleLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10" opacity=".5"/><path stroke-linecap="round" stroke-linejoin="round" d="m8 12l2-1.5L8 9m8 3l-2-1.5L16 9m-1 7l-1-1l-1 1l-1-1l-1 1l-1-1l-1 1"/></g>`),
		g.Group(children),
	)
}
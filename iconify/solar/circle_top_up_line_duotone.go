package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleTopUpLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m13 11l9-9m0 0h-5.344M22 2v5.344"/><path d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10" opacity=".5"/></g>`),
		g.Group(children),
	)
}
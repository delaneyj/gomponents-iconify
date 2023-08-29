package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleTopDownLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m21 3l-9 9m0 0h5.344M12 12V6.656"/><path d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10"/></g>`),
		g.Group(children),
	)
}
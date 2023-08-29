package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextCrossLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12 3H8c-1.886 0-2.828 0-3.414.586C4 4.172 4 5.114 4 7v.95M12 3h4c1.886 0 2.828 0 3.414.586C20 4.172 20 5.114 20 7v.95M12 3v6m0 12v-6m-5 6h10"/><path d="M4 12h16"/></g>`),
		g.Group(children),
	)
}
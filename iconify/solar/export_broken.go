package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExportBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M12 20a8 8 0 0 1-8-8m16 0a7.985 7.985 0 0 1-3 6.245"/><path stroke-linejoin="round" d="M12 14V4m0 0l3 3m-3-3L9 7"/></g>`),
		g.Group(children),
	)
}
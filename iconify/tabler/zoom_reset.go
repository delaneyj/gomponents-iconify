package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomReset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m21 21l-6-6M3.268 12.043A7.017 7.017 0 0 0 9.902 17a7.012 7.012 0 0 0 7.043-6.131a7 7 0 0 0-5.314-7.672A7.021 7.021 0 0 0 3.39 7.6"/><path d="M3 4v4h4"/></g>`),
		g.Group(children),
	)
}
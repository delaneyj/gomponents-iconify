package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceMask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 14.5h-.222C3.243 14.5 2 13.38 2 12s1.243-2.5 2.778-2.5H5m14 5h.222C20.756 14.5 22 13.38 22 12s-1.244-2.5-2.778-2.5H19M9 10h6m-6 4h6"/><path d="m12.55 18.843l5-1.429A2 2 0 0 0 19 15.491V8.51a2 2 0 0 0-1.45-1.923l-5-1.429a2 2 0 0 0-1.1 0l-5 1.429A2 2 0 0 0 5 8.509v6.982a2 2 0 0 0 1.45 1.923l5 1.429a2 2 0 0 0 1.1 0z"/></g>`),
		g.Group(children),
	)
}
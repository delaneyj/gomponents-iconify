package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceMaskOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 14.5h-.222C3.243 14.5 2 13.38 2 12s1.243-2.5 2.778-2.5H5m14 5h.222C20.756 14.5 22 13.38 22 12s-1.244-2.5-2.778-2.5H19M9 10h1m4 0h1m-6 4h5"/><path d="M19 15V8.51a2 2 0 0 0-1.45-1.923l-5-1.429a2 2 0 0 0-1.1 0l-1.788.511m-3.118.891l-.094.027A2 2 0 0 0 5 8.509v6.982a2 2 0 0 0 1.45 1.923l5 1.429a2 2 0 0 0 1.1 0l4.899-1.4M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
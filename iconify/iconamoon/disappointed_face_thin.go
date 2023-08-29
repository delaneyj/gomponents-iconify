package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisappointedFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M10 15.535A3.981 3.981 0 0 1 12 15a3.98 3.98 0 0 1 2 .535m3-5.035l-3-1m-4 0l-3 1"/></g>`),
		g.Group(children),
	)
}
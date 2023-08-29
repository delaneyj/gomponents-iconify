package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterdropBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M3.447 16A8.98 8.98 0 0 1 3 13.193v-.265C3 8.317 5.729 4.163 9.903 2.421a5.434 5.434 0 0 1 4.194 0C18.272 4.163 21 8.317 21 12.928v.265C21 18.057 17.145 22 12.389 22h-.777A8.467 8.467 0 0 1 6 19.874"/><path d="M7.615 10.724c.634-2.006 2.016-3.636 3.77-4.448"/></g>`),
		g.Group(children),
	)
}
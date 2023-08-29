package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarBriefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 0h1v3H3V0zm8 0h1v3h-1V0z"/><path fill="currentColor" d="M13 1v3h-3V1H5v3H2V1H0v14h5v-1H1V6h13v3h1V1z"/><path fill="currentColor" d="M13 10V8H9v2H6v6h10v-6h-3zm-3-1h2v1h-2V9z"/>`),
		g.Group(children),
	)
}
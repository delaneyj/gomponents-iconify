package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarageLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M22 22H2m1 0V11.347a3 3 0 0 1 1.007-2.242l6-5.333a3 3 0 0 1 3.986 0l6 5.333A3 3 0 0 1 21 11.347V22M10 9h4m-5 6.5h6m-6 3h6"/><path d="M18 22v-6c0-1.886 0-2.828-.586-3.414C16.828 12 15.886 12 14 12h-4c-1.886 0-2.828 0-3.414.586C6 13.172 6 14.114 6 16v6"/></g>`),
		g.Group(children),
	)
}
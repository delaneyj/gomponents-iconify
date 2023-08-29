package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingsBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M22 22H2m19 0v-9m-9.996-8c.018-1.24.11-1.943.582-2.414C12.172 2 13.114 2 15 2h2c1.886 0 2.828 0 3.414.586C21 3.172 21 4.114 21 6v3"/><path d="M15 22v-6M3 22v-9m0-4c0-1.886 0-2.828.586-3.414C4.172 5 5.114 5 7 5h4c1.886 0 2.828 0 3.414.586C15 6.172 15 7.114 15 9v3M9 22v-3M6 8h6m-6 3h1m5 0H9.5M6 14h6"/></g>`),
		g.Group(children),
	)
}
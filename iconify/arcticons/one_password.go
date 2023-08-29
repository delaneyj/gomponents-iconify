package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OnePassword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.083 17.28a7.863 7.863 0 0 1 0 13.44m-8.166 0a7.863 7.863 0 0 1 0-13.44m6.15-6.85h-4.134a2.017 2.017 0 0 0-2.016 2.016v6.36c2.358 1.281 2.736 2.562 0 3.843v12.925a2.017 2.017 0 0 0 2.016 2.015h4.134a2.017 2.017 0 0 0 2.016-2.015v-6.361c-2.358-1.281-2.736-2.562 0-3.842V12.446a2.017 2.017 0 0 0-2.016-2.016Z"/>`),
		g.Group(children),
	)
}
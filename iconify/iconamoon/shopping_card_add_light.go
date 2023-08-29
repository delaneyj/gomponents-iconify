package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCardAddLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 2.25a.75.75 0 0 0 0 1.5v-1.5ZM5 3l.748-.058A.75.75 0 0 0 5 2.25V3Zm16 3l.745.083A.75.75 0 0 0 21 5.25V6ZM5.23 6l-.747.058L5.231 6Zm13.109 9.119l.053.748l-.053-.748Zm-10.355.74l-.053-.749l.053.748ZM3 3.75h2v-1.5H3v1.5Zm5.037 12.856l10.355-.74l-.107-1.495l-10.354.74l.106 1.495Zm12.892-3.179l.816-7.344l-1.49-.166l-.816 7.345l1.49.165ZM4.252 3.057l.231 3l1.496-.115l-.231-3l-1.496.116Zm.231 3l.617 8.017l1.495-.115l-.616-8.017l-1.496.116ZM21 5.25H5.23v1.5H21v-1.5Zm-2.608 10.617a2.75 2.75 0 0 0 2.537-2.44l-1.49-.165a1.25 1.25 0 0 1-1.154 1.109l.107 1.496ZM7.931 15.11a1.25 1.25 0 0 1-1.336-1.15l-1.495.114a2.75 2.75 0 0 0 2.937 2.532l-.106-1.496Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 9.5v3M11.5 11h3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2.25" d="M17.5 20.5h.01v.01h-.01zm-9 0h.01v.01H8.5z"/></g>`),
		g.Group(children),
	)
}
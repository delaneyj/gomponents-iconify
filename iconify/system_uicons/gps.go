package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" transform="translate(2 2)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 14.5c3.329 0 6-2.645 6-5.973c0-3.329-2.671-6.027-6-6.027s-6 2.698-6 6.027c0 3.328 2.671 5.973 6 5.973z"/><circle cx="8.5" cy="8.5" r="3.5" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 8.5h2m12 0h2m-8-8v2m0 12v2"/></g>`),
		g.Group(children),
	)
}
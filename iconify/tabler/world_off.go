package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorldOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.657 5.615a9 9 0 1 0 12.717 12.739m1.672-2.322A9 9 0 0 0 7.98 3.948M3.6 9H9m4 0h7.4M3.6 15H15m4 0h1.4"/><path d="M11.5 3a17.001 17.001 0 0 0-1.493 3.022M9.16 9.167c-.68 4.027.1 8.244 2.34 11.833m1-18a16.982 16.982 0 0 1 2.549 8.005m-.207 3.818A16.979 16.979 0 0 1 12.5 21M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
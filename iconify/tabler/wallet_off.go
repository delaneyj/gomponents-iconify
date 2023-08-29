package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 8V5a1 1 0 0 0-1-1H8m-3.413.584A2 2 0 0 0 6 8h2m4 0h6a1 1 0 0 1 1 1v3"/><path d="M19 19a1 1 0 0 1-1 1H6a2 2 0 0 1-2-2V6"/><path d="M16 12h4v4m-4 0a2 2 0 0 1-2-2M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
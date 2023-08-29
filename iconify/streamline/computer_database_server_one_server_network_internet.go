package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDatabaseServerOneServerNetworkInternet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="6" x=".5" y="1" rx="1"/><circle cx="3.5" cy="4" r=".5"/><path d="M7.5 4H11M1.5 7a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1"/><circle cx="3.5" cy="10" r=".5"/><path d="M7.5 10H11"/></g>`),
		g.Group(children),
	)
}
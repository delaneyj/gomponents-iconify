package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDatabaseServerThreeServerNetworkInternet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="6" x=".5" y="3" rx="1"/><circle cx="3.5" cy="6" r=".5"/><path d="M6.5 6h2M7 9v4.5m-5 0h10"/></g>`),
		g.Group(children),
	)
}
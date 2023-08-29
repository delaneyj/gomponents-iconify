package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerConnectionNetworkNetworkServerInternetEthernet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 10.5v3m-5 0h10"/><circle cx="7" cy="5.5" r="5"/><path d="M2 5.5h10m-5 5c3-3.42 3-6.76 0-10c-2.94 3.12-3 6.44 0 10Z"/></g>`),
		g.Group(children),
	)
}
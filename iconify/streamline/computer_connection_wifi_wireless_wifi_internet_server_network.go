package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerConnectionWifiWirelessWifiInternetServerNetwork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="10.94" r="1.31"/><path d="M4.53 8a3.49 3.49 0 0 1 5 0M2.36 6.31a6.55 6.55 0 0 1 9.29 0"/><path d="M.5 4.45a9.19 9.19 0 0 1 13 0"/></g>`),
		g.Group(children),
	)
}
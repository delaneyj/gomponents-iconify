package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerConnectionSignalLoadingBracketLoadingInternetAngleSignalServerNetworkConnecting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3 9.5L.5 7L3 4.5m8 5L13.5 7L11 4.5"/><circle cx="9" cy="7" r=".5"/><circle cx="5" cy="7" r=".5"/></g>`),
		g.Group(children),
	)
}
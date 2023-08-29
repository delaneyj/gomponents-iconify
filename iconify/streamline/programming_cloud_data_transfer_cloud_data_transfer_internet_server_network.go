package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingCloudDataTransferCloudDataTransferInternetServerNetwork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12 8.1a3 3 0 1 0-3.65-4.69A4 4 0 1 0 1.38 7M3.5 9.5l2-2v6"/><path d="m10.5 11.5l-2 2v-6"/></g>`),
		g.Group(children),
	)
}
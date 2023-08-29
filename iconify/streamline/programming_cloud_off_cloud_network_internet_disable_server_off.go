package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingCloudOffCloudNetworkInternetDisableServerOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.17 3.17a4.08 4.08 0 0 1 6.18 2.24a3 3 0 1 1 2.24 5.09l-1 .05M.75 5.1A4 4 0 0 0 .5 6.5a4 4 0 0 0 4 4H6m-5.5-9l11 11"/>`),
		g.Group(children),
	)
}
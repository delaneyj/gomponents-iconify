package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingCloudAddCloudNetworkInternetAddServerPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 8.5a3 3 0 1 0-2.15-5.09A4 4 0 1 0 4.5 8.5M10 11H5m2.5-2.5v5"/>`),
		g.Group(children),
	)
}
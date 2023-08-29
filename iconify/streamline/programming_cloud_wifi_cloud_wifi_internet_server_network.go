package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingCloudWifiCloudWifiInternetServerNetwork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 13.5a3 3 0 1 0-2.15-5.09A4 4 0 1 0 4.5 13.5Zm-.02-12.3a9 9 0 0 0-7 0m5.7 2.16a6.78 6.78 0 0 0-4.36 0"/>`),
		g.Group(children),
	)
}
package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopCloudAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 10H10a3 3 0 0 0 1.07-5.8a4 4 0 0 0-7.48 1A2.5 2.5 0 0 0 4.5 10Zm0-3a1 1 0 0 0 1-1a2 2 0 0 1 3.89-.64a1 1 0 0 0 .78.66A1 1 0 0 1 11 7a1 1 0 0 1-1 1H4.5a.5.5 0 0 1 0-1ZM19 2h-5a1 1 0 0 0 0 2h5a1 1 0 0 1 1 1v7H3a1 1 0 0 0-1 1v2a3 3 0 0 0 3 3h3v2H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2h-3v-2h3a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3Zm-5 18h-4v-2h4Zm6-5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-1h16Z"/>`),
		g.Group(children),
	)
}
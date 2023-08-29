package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentWindows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23H4M18 1v5h5m-8.75 5.5l-6 7m0-7l6 7M20.5 12h-4v6h4m-1-3h-3M7 12H3v6h4m-1-3H3"/>`),
		g.Group(children),
	)
}
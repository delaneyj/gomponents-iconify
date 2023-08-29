package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskVisible(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 1H7v2H3v20h7v-2H5V5h2v2h10V5h2v8h2V3h-4V1Zm-2 4H9V3h6v2Zm.75 15v-2h2.5v2h-2.5Z"/><path fill="currentColor" d="M17.002 23.5c4.419 0 6.09-4.5 6.09-4.5s-1.673-4.5-6.09-4.5c-4.416 0-6.09 4.5-6.09 4.5s1.672 4.5 6.09 4.5Zm-.002-2c-2.615 0-3.87-2.5-3.87-2.5s1.25-2.5 3.87-2.5c2.621 0 3.87 2.5 3.87 2.5s-1.254 2.5-3.87 2.5Z"/>`),
		g.Group(children),
	)
}
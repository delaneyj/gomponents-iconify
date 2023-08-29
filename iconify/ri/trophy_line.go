package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrophyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.005 16.941v2.062h5v2h-12v-2h5V16.94a8.001 8.001 0 0 1-7-7.938v-6h16v6a8.001 8.001 0 0 1-7 7.938Zm-7-11.938v4a6 6 0 0 0 12 0v-4h-12Zm-5 0h2v4h-2v-4Zm20 0h2v4h-2v-4Z"/>`),
		g.Group(children),
	)
}
package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloggerAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 23H4a3.003 3.003 0 0 1-3-3V4a3.003 3.003 0 0 1 3-3h16a3.003 3.003 0 0 1 3 3v16a3.003 3.003 0 0 1-3 3z" opacity=".5"/><path fill="currentColor" d="M16.003 10.002H16V9a4.004 4.004 0 0 0-4-4h-2a5.006 5.006 0 0 0-5 5v4a5.006 5.006 0 0 0 5 5h4a5.006 5.006 0 0 0 5-5v-1a3 3 0 0 0-2.997-2.998zM10 9h1a1 1 0 1 1 0 2h-1a1 1 0 1 1 0-2zm4 6h-4a1 1 0 1 1 0-2h4a1 1 0 1 1 0 2z"/>`),
		g.Group(children),
	)
}
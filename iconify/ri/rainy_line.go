package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RainyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 18v-2h1a4 4 0 1 0-2.157-7.37A6 6 0 1 0 8 15.917v2.022A8.001 8.001 0 0 1 9 2a7.998 7.998 0 0 1 6.98 4.087A6 6 0 1 1 17 18h-1Zm-5.768.732L12 16.965l1.768 1.767a2.5 2.5 0 1 1-3.536 0Z"/>`),
		g.Group(children),
	)
}
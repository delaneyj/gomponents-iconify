package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Palace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 2v1.126c1.725.444 3 2.01 3 3.874h3v2h-1v1h4v2h-1v8h1v2H2v-2h1v-8H2v-2h4V9H5V7h3a4.002 4.002 0 0 1 3-3.874V2h2ZM8 9v1h8V9H8Zm6-2a2 2 0 1 0-4 0h4Zm-9 5v8h3v-2a4 4 0 0 1 8 0v2h3v-8H5Zm9 8v-2a2 2 0 1 0-4 0v2h4Z"/>`),
		g.Group(children),
	)
}
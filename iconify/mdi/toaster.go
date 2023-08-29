package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toaster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 11a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2H2v2h1v7h18v-7h1v-2m-5 4a2 2 0 1 1 2-2a2 2 0 0 1-2 2m1-7H6c.33-2.25 2.88-4 6-4s5.63 1.75 6 4Z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerSocketUk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H4m0 2h16v16H4V4m7 3v4h2V7h-2m-5 7.75V17h3.5v-2.25H6m8.5 0V17H18v-2.25h-3.5Z"/>`),
		g.Group(children),
	)
}
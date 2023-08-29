package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PostJp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.5 4a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2h-10a1 1 0 0 1-1-1zm11 2.5h-10a1 1 0 0 0 0 2h4V13a1 1 0 1 0 2 0V8.5h4a1 1 0 1 0 0-2z"/>`),
		g.Group(children),
	)
}
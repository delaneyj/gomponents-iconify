package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Activity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2H2v20h20V2Zm-2 2v7h-2.554l-2.021 3.233l-5.865-7.82L5.546 11H4V4h16ZM4 13h2.454L9.44 9.587l6.135 8.18L18.555 13H20v7H4v-7Z"/>`),
		g.Group(children),
	)
}
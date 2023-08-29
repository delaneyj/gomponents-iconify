package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-opacity=".5" d="M9 13h6v4H9v-4Z"/><path d="M6 7h12v4H6V7Z"/></g>`),
		g.Group(children),
	)
}
package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-opacity=".5" d="M8 13h6v4H8v-4Z"/><path d="M6 6H4v12h2V6Zm14 1H8v4h12V7Z"/></g>`),
		g.Group(children),
	)
}
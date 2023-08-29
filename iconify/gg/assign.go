package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Assign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 6h4V4H4v6h2V6Zm4 12H6v-4H4v6h6v-2Zm4-12h4v4h2V4h-6v2Zm0 12h4v-4h2v6h-6v-2Zm-2-9.5a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Z"/>`),
		g.Group(children),
	)
}
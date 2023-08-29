package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m12 16l5 2V8l-5-2l-5 2v10l5-2Zm-3-.954l3-1.2l3 1.2V9.354l-3-1.2l-3 1.2v5.692Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
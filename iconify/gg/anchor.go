package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 6a3.001 3.001 0 0 1-2 2.83v8.044c1.725-.444 3-2.01 3-3.874h2a6.002 6.002 0 0 1-5 5.917V20a1 1 0 1 1-2 0v-1.083A6.002 6.002 0 0 1 6 13h2a4.002 4.002 0 0 0 3 3.874V8.829A3.001 3.001 0 1 1 15 6Zm-3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
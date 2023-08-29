package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 8v5a5 5 0 0 0 10 0V8H7ZM5 4h14v9a7 7 0 1 1-14 0V4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExtensionAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 5v14h8v-6h6V5H5Zm6 2H7v4h4V7Zm0 6H7v4h4v-4Zm2-2h4V7h-4v4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
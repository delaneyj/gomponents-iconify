package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 3H6a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h11a2 2 0 0 0 1.41-.58A2 2 0 0 0 19 20V2ZM6 6h8v2H6Zm4 14H8v-2h2Zm0-6h4v2h-4Zm7 6h-6v-2h6Zm0-4h-2v-2h2Zm0-4h-5v-2h5Zm0-4h-2V6h2Z"/>`),
		g.Group(children),
	)
}
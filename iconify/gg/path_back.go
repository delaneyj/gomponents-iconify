package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PathBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 14H4V4h10v5h5v10H9v-5ZM6 6h6v6H6V6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
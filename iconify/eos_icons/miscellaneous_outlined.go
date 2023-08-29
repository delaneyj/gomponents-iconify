package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiscellaneousOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.36 15L18 21H6l1.64-6h8.72M18 13H6l-2 8a2.185 2.185 0 0 0 2 2h12a2.158 2.158 0 0 0 2-2l-2-8Zm4-3H2v2h20v-2zM10 8a1 1 0 0 1 1 1h2a1 1 0 0 1 2 0h2a5.11 5.11 0 0 0-2.286-3.196A6.303 6.303 0 0 0 15 3.835C15 2.27 14.552 1 14 1s-1 1.27-1 2.835a7.115 7.115 0 0 0 .115 1.301a4.626 4.626 0 0 0-2.234.001A7.094 7.094 0 0 0 11 3.835C11 2.27 10.552 1 10 1S9 2.27 9 3.835a6.31 6.31 0 0 0 .283 1.971A5.11 5.11 0 0 0 7 9h2a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}
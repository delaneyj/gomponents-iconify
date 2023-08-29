package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 33V15M10 9h28v6H10zM8 32l6-7h19.974L40 32M4 33h8v8H4zm16 0h8v8h-8zm16 0h8v8h-8z"/>`),
		g.Group(children),
	)
}
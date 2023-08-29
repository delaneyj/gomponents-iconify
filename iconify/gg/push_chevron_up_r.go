package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushChevronUpR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 21h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2Zm-4-2a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v14Zm8.172-2.288l-1.415-1.415L12 11.055l4.243 4.242l-1.415 1.415L12 13.883l-2.828 2.829ZM8 10h8V8H8v2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
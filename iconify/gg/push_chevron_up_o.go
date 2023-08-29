package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushChevronUpO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18Zm0-2C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1Zm-4 9V8h8v2H8Zm8.243 5.297l-1.414 1.415L12 13.883l-2.828 2.829l-1.415-1.415L12 11.055l4.243 4.242Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
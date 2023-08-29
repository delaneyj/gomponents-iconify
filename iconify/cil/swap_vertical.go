package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M384 433.373V160h-32v274.51l-69.823-69.823l-22.627 22.626l107.882 107.883l107.881-107.883l-22.626-22.626L384 433.373zM159.432 17.372L51.55 125.255l22.627 22.627L144 78.059V352h32V79.195l68.687 68.687l22.626-22.627L159.432 17.372z"/>`),
		g.Group(children),
	)
}
package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoAdobePhotoshop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 13v5H6V6h3.5a3.5 3.5 0 1 1 0 7H8Zm0-5v3h1.5a1.5 1.5 0 0 0 0-3H8Zm5.5 4.453A2.453 2.453 0 0 1 15.953 10H18v2h-2.047a.453.453 0 0 0-.143.883l1.013.337a2.453 2.453 0 0 1-.776 4.78H13.5v-2h2.547a.453.453 0 0 0 .143-.883l-1.013-.337a2.453 2.453 0 0 1-1.677-2.327Z"/><path fill="currentColor" d="M22 2H2v20h20V2ZM4 20V4h16v16H4Z"/>`),
		g.Group(children),
	)
}
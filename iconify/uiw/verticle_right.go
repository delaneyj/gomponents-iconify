package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.606 9.256L6.614 2.203c-.367-.292-.704-.292-1.012 0c-.308.291-.308.616 0 .974l6.626 6.678l-6.992 6.876c-.313.304-.321.625-.023.966c.299.34.64.347 1.023.02l7.37-7.245v6.836c0 .461.233.692.697.692c.465 0 .697-.23.697-.692V2.692c0-.466-.232-.697-.697-.692c-.464.005-.697.236-.697.692v6.564Z"/>`),
		g.Group(children),
	)
}
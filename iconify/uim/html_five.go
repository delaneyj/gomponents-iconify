package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HtmlFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.182 2l1.605 18l7.202 2l7.222-2.002L20.818 2Zm14.143 5.887H8.877l.202 2.261h8.045l-.606 6.778L12 18.178l-.01.003l-4.523-1.255l-.309-3.466h2.216l.158 1.76l2.458.664h.002l2.463-.665l.256-2.863H7.06L6.464 5.68h11.059Z"/><path fill="currentColor" d="M17.325 7.887H8.877l.202 2.261h8.045l-.606 6.778L12 18.178l-.01.003l-4.523-1.255l-.309-3.466h2.216l.158 1.76l2.458.664h.002l2.463-.665l.256-2.863H7.06L6.464 5.68h11.059Z" opacity=".25"/>`),
		g.Group(children),
	)
}
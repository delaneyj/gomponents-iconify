package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Milk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.85 1H6.15v3.59L4 9.49V23h16V9.49l-2.15-4.9V1Zm-.918 6.476L18 9.91V21h-1.95V9.878l.882-2.402ZM14.05 10.7V21H6V10.7h8.05Zm-7.52-2l1.273-2.9h7.614l-1.065 2.9H6.531Zm1.62-4.9V3h7.7v.8h-7.7ZM13 13h-1.426l-1.576 1.584L8.415 13H7v6h2v-2.586l1.002 1.002l.998-1.003V19h2v-6Z"/>`),
		g.Group(children),
	)
}
package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dividers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 1v1.232a4.359 4.359 0 0 1 3.375 4.239a4.32 4.32 0 0 1-.678 2.33L18.966 15H22v2h-1.979l2.33 4.418l-1.768.933L17.76 17H13v2h-2v-2H6.24l-2.822 5.351l-1.769-.933L3.98 17H2v-2h3.035l3.27-6.2a4.32 4.32 0 0 1-.68-2.33A4.359 4.359 0 0 1 11 2.233V1h2Zm-3.193 9.238L7.296 15H11v-2h2v2h3.705l-2.511-4.762a4.375 4.375 0 0 1-2.194.585a4.375 4.375 0 0 1-2.193-.585ZM12 4.118A2.364 2.364 0 0 0 9.625 6.47c0 .6.227 1.148.603 1.566A2.376 2.376 0 0 0 12 8.824c.706 0 1.338-.304 1.773-.787a2.33 2.33 0 0 0 .602-1.566A2.364 2.364 0 0 0 12 4.118Z"/>`),
		g.Group(children),
	)
}
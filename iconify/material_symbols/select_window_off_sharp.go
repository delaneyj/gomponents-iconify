package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectWindowOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.475 23.3L10.175 13H4v7h12v-4.025l2 2V22H2V9h4v-.175L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM18 15.125L15.875 13l-4-4H18v4h2V6H8.875L6.15 3.275V2H22v13h-4v.125Z"/>`),
		g.Group(children),
	)
}
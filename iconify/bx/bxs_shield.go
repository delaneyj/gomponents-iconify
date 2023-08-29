package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsShield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M20.496 6.106l-7.973-4a.997.997 0 0 0-.895-.002l-8.027 4c-.297.15-.502.437-.544.767c-.013.097-1.145 9.741 8.541 15.008a.995.995 0 0 0 .969-.009c9.307-5.259 8.514-14.573 8.476-14.967a1 1 0 0 0-.547-.797z" fill="currentColor"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audiometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" d="M6.63 36.5h34.74a.12.12 0 0 1 .13.12v6.76a.12.12 0 0 1-.13.12H6.63a.12.12 0 0 1-.13-.12v-6.76a.12.12 0 0 1 .13-.12Zm0-32h34.74a.12.12 0 0 1 .13.12v6.76a.12.12 0 0 1-.13.12H6.63a.12.12 0 0 1-.13-.12V4.62a.12.12 0 0 1 .13-.12Zm0 21.33h34.74a.13.13 0 0 1 .13.12v6.76a.12.12 0 0 1-.13.12H6.63a.12.12 0 0 1-.13-.12V26a.13.13 0 0 1 .13-.17Zm0-10.66h34.74a.12.12 0 0 1 .13.12v6.76a.13.13 0 0 1-.13.12H6.63a.13.13 0 0 1-.13-.12v-6.76a.12.12 0 0 1 .13-.12Z"/>`),
		g.Group(children),
	)
}
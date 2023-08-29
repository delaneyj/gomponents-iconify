package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoyalTsd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.262 16.595L11.585 35.792h25.528l6.149-19.197zm-13.026 7.893l-6.214-12.28l-12.437 23.584"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.585 35.792l-7.323-18.72l12.955 8.039"/>`),
		g.Group(children),
	)
}
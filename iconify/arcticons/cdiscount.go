package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cdiscount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.28 13.418A21.506 21.506 0 0 0 24 45.5a21.717 21.717 0 0 0 2.657-.163M42.72 34.583A21.506 21.506 0 0 0 24 2.5a21.714 21.714 0 0 0-2.658.163"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.046 15.954a11.378 11.378 0 1 0 0 16.092"/>`),
		g.Group(children),
	)
}
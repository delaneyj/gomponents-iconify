package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Solarman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.872 4.5S9.99 18.065 9.99 29.49a14.01 14.01 0 0 0 28.018 0c0-8.149-9.187-14.32-9.187-14.32l-8.38 12.677h7.706l-6.314 9.55"/>`),
		g.Group(children),
	)
}
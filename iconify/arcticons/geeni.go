package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Geeni(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.982 20.136a7.76 7.76 0 1 1 1.04-3.88m7.956-.001a7.76 7.76 0 1 1 1.04 3.88m-1.04-3.88l15.521-.04m-23.477.161H4.502m5.894 18.265c8.564 6.621 17.654 6.35 27.207 0"/>`),
		g.Group(children),
	)
}
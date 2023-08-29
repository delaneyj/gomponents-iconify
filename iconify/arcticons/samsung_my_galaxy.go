package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungMyGalaxy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.684 10.506C14.179 13.176.678 21.32 4.457 30.263c6.287 14.882 41.963 1.683 38.068-9.176c-1.809-5.042-12.85-1.397-21.974 1.178"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plthub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 15.91V3.5m0 41V32.09m3.37-11.46L38.5 9.5m-29 29l11.13-11.13M32.09 24H44.5m-41 0h12.41m13.81 5.72l8.78 8.78m-29-29l8.78 8.78"/><circle cx="24" cy="24" r="4.76" fill="none" stroke="currentColor"/>`),
		g.Group(children),
	)
}
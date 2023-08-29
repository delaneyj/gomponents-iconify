package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryChargingVerticalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M149.1 132.85a6 6 0 0 1 .27 5.83l-16 32a6 6 0 1 1-10.74-5.36L134.29 142H112a6 6 0 0 1-5.37-8.68l16-32a6 6 0 0 1 10.74 5.36L121.71 130H144a6 6 0 0 1 5.1 2.85ZM96 14h64a6 6 0 0 0 0-12H96a6 6 0 0 0 0 12Zm102 42v168a22 22 0 0 1-22 22H80a22 22 0 0 1-22-22V56a22 22 0 0 1 22-22h96a22 22 0 0 1 22 22Zm-12 0a10 10 0 0 0-10-10H80a10 10 0 0 0-10 10v168a10 10 0 0 0 10 10h96a10 10 0 0 0 10-10Z"/>`),
		g.Group(children),
	)
}
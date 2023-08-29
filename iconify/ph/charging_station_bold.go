package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChargingStationBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M137.94 121.27a12 12 0 0 1 1.2 11.19l-16 40A12 12 0 0 1 112 180a11.86 11.86 0 0 1-4.46-.86a12 12 0 0 1-6.68-15.6l9.42-23.54H96a12 12 0 0 1-11.14-16.46l16-40a12 12 0 0 1 22.28 8.92L113.72 116H128a12 12 0 0 1 9.94 5.27ZM256 86.63V166a30 30 0 0 1-60 0v-38a4 4 0 0 0-4-4h-12v80h12a12 12 0 0 1 0 24H32a12 12 0 0 1 0-24h12V56a28 28 0 0 1 28-28h80a28 28 0 0 1 28 28v44h12a28 28 0 0 1 28 28v38a6 6 0 0 0 12 0V86.63a4 4 0 0 0-1.17-2.83l-19.32-19.31a12 12 0 0 1 17-17l19.29 19.34a27.81 27.81 0 0 1 8.2 19.8ZM156 204V56a4 4 0 0 0-4-4H72a4 4 0 0 0-4 4v148Z"/>`),
		g.Group(children),
	)
}
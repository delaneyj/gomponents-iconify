package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryChargingBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M196 52H28A28 28 0 0 0 0 80v96a28 28 0 0 0 28 28h168a28 28 0 0 0 28-28V80a28 28 0 0 0-28-28Zm4 124a4 4 0 0 1-4 4H28a4 4 0 0 1-4-4V80a4 4 0 0 1 4-4h168a4 4 0 0 1 4 4Zm56-72v48a12 12 0 0 1-24 0v-48a12 12 0 0 1 24 0Zm-113.62 18a12 12 0 0 1 0 12l-16 28a12 12 0 1 1-20.84-11.9l5.78-10.1H100a12 12 0 0 1-10.42-18l16-28a12 12 0 0 1 20.84 12l-5.74 10H132a12 12 0 0 1 10.38 6Z"/>`),
		g.Group(children),
	)
}
package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirZigzag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M185.901 21.667c87.765 81.467 140 169.208 177.287 267.486c-87.698-104.24-120.874-125.43-246.298-179.823c91.563 72.21 109.218 91.017 162.839 176.715c-70.043-52.221-166.377-103.052-258.966-130.635v85.401c95.446 24.54 244.61 97.1 311.842 175.215c-.625-25.575-3.985-51.147-14.516-76.722L426.95 451.276c-6.545-34.005-20.51-67.185-41.99-99.53l114.046 116.636C461.763 292.837 343.066 104.34 185.903 21.667Z"/>`),
		g.Group(children),
	)
}
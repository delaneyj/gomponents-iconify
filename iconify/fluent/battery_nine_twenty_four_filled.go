package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryNineTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 9a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1H6ZM2 9a3 3 0 0 1 3-3h12.5a3 3 0 0 1 3 3v1h.5a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-.5v1a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V9Z"/>`),
		g.Group(children),
	)
}
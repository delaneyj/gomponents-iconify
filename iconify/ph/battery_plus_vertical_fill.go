package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryPlusVerticalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M88 8a8 8 0 0 1 8-8h64a8 8 0 0 1 0 16H96a8 8 0 0 1-8-8Zm112 48v168a24 24 0 0 1-24 24H80a24 24 0 0 1-24-24V56a24 24 0 0 1 24-24h96a24 24 0 0 1 24 24Zm-36 84a8 8 0 0 0-8-8h-20v-20a8 8 0 0 0-16 0v20h-20a8 8 0 0 0 0 16h20v20a8 8 0 0 0 16 0v-20h20a8 8 0 0 0 8-8Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M256 96v64a8 8 0 0 1-16 0V96a8 8 0 0 1 16 0Zm-32-16v96a24 24 0 0 1-24 24H32a24 24 0 0 1-24-24V80a24 24 0 0 1 24-24h168a24 24 0 0 1 24 24Zm-16 0a8 8 0 0 0-8-8H32a8 8 0 0 0-8 8v96a8 8 0 0 0 8 8h168a8 8 0 0 0 8-8Zm-92 52a8 8 0 0 0 8-8V96a8 8 0 0 0-16 0v28a8 8 0 0 0 8 8Zm0 12a12 12 0 1 0 12 12a12 12 0 0 0-12-12Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadHorizonBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M237.88 202.46a12 12 0 0 1-16.34-4.58L153 76h-13v4a12 12 0 0 1-24 0v-4h-13L34.46 197.88a12 12 0 1 1-20.92-11.76L75.48 76H24a12 12 0 0 1 0-24h208a12 12 0 0 1 0 24h-51.48l61.94 110.12a12 12 0 0 1-4.58 16.34ZM128 108a12 12 0 0 0-12 12v16a12 12 0 0 0 24 0v-16a12 12 0 0 0-12-12Zm0 56a12 12 0 0 0-12 12v16a12 12 0 0 0 24 0v-16a12 12 0 0 0-12-12Z"/>`),
		g.Group(children),
	)
}
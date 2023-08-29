package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCreamFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 89.37V88a80 80 0 0 0-160 0v1.37A24 24 0 0 0 56 136h3.36l61.69 108a8 8 0 0 0 13.9 0l61.69-108H200a24 24 0 0 0 8-46.63Zm-80 134.51L77.79 136h19.28l40.57 71Zm18.86-33L115.5 136h19.29l21.71 38Zm18.85-33L153.21 136h25Z"/>`),
		g.Group(children),
	)
}
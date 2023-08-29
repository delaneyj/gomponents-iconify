package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FadersBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M140 124v92a12 12 0 0 1-24 0v-92a12 12 0 0 1 24 0Zm60 68a12 12 0 0 0-12 12v12a12 12 0 0 0 24 0v-12a12 12 0 0 0-12-12Zm24-40h-12V40a12 12 0 0 0-24 0v112h-12a12 12 0 0 0 0 24h48a12 12 0 0 0 0-24Zm-168 8a12 12 0 0 0-12 12v44a12 12 0 0 0 24 0v-44a12 12 0 0 0-12-12Zm24-40H68V40a12 12 0 0 0-24 0v80H32a12 12 0 0 0 0 24h48a12 12 0 0 0 0-24Zm72-48h-12V40a12 12 0 0 0-24 0v32h-12a12 12 0 0 0 0 24h48a12 12 0 0 0 0-24Z"/>`),
		g.Group(children),
	)
}
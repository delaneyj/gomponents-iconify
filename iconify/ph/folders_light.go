package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoldersLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 66h-69.33a2 2 0 0 1-1.2-.4l-27.74-20.8a14 14 0 0 0-8.4-2.8H72a14 14 0 0 0-14 14v18H40a14 14 0 0 0-14 14v112a14 14 0 0 0 14 14h152.89A13.12 13.12 0 0 0 206 200.89V182h18.89A13.12 13.12 0 0 0 238 168.89V80a14 14 0 0 0-14-14Zm-30 134.89a1.11 1.11 0 0 1-1.11 1.11H40a2 2 0 0 1-2-2V88a2 2 0 0 1 2-2h45.33a2 2 0 0 1 1.2.4l27.74 20.8a14 14 0 0 0 8.4 2.8H192a2 2 0 0 1 2 2Zm32-32a1.11 1.11 0 0 1-1.11 1.11H206v-58a14 14 0 0 0-14-14h-69.33a2 2 0 0 1-1.2-.4L93.73 76.8a14 14 0 0 0-8.4-2.8H70V56a2 2 0 0 1 2-2h45.33a2 2 0 0 1 1.2.4l27.74 20.8a14 14 0 0 0 8.4 2.8H224a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}
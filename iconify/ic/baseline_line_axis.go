package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineLineAxis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 7.43l-1.41-1.41l-4.03 4.53L9.5 4L2 11.51l1.5 1.5l6.14-6.15l5.59 5.18l-1.73 1.95l-4-4L2 17.5L3.5 19l6-6.01l4 4l3.19-3.59l3.9 3.61L22 15.6l-3.98-3.7z"/>`),
		g.Group(children),
	)
}
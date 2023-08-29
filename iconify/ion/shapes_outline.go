package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" d="M336 320H32L184 48l152 272z"/><path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" d="M265.32 194.51A144 144 0 1 1 192 320"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.475 23.3L16 18.8l-4 4L1.2 12l5.425-5.425L12.25 12.2V15L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM8 15h2v-3h2.025l-2-2H9q-.425 0-.713.288T8 11v4Zm10.85.95l-3.4-3.4L17 11l-3.5-3.5V10h-.6L8.05 5.15L12 1.2L22.8 12l-3.95 3.95Z"/>`),
		g.Group(children),
	)
}
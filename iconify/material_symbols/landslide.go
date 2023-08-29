package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Landslide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22v-4.25l4 1.3l11.025-3.675L22 22H2Zm4-5.05L2 15.6v-1.85l4 1.3l6.9-2.3l2.55 1.025L6 16.95ZM18.5 14l4.5-2V8l-4.5-1L16 9v3l2.5 2ZM6 12.95L2 11.6V8h6l2.575 3.425L6 12.95ZM12 8l5-2V1l-5-1l-3 2v4l3 2Z"/>`),
		g.Group(children),
	)
}
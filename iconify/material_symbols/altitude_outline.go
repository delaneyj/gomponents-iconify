package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AltitudeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 12V7.8l-1.6 1.6L15 8l4-4l4 4l-1.4 1.425l-1.6-1.6V12h-2ZM1 22l6-8l4.5 6H19l-5-6.65l-2.5 3.3L10.25 15L14 10l9 12H1Zm10.5-2Z"/>`),
		g.Group(children),
	)
}
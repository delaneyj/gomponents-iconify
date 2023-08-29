package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarsStrokeV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 1.594l-.719.687l-6 6L10.72 9.72L15 5.438V10h-3v2h3v2.063c-3.934.5-7 3.87-7 7.937c0 4.406 3.594 8 8 8c4.406 0 8-3.594 8-8c0-4.066-3.066-7.438-7-7.938V12h3v-2h-3V5.437l4.281 4.282L22.72 8.28l-6-6zM16 16c3.324 0 6 2.676 6 6s-2.676 6-6 6s-6-2.676-6-6s2.676-6 6-6z"/>`),
		g.Group(children),
	)
}
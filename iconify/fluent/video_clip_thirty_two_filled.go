package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoClipThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 4A4.5 4.5 0 0 0 2 8.5v15A4.5 4.5 0 0 0 6.5 28h19a4.5 4.5 0 0 0 4.5-4.5v-15A4.5 4.5 0 0 0 25.5 4h-19Zm5.5 8.001a1 1 0 0 1 1.47-.882l7.498 3.999a1 1 0 0 1 0 1.764l-7.497 3.999A1 1 0 0 1 12 19.999V12Z"/>`),
		g.Group(children),
	)
}
package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m8.65 1.158l-5.485 5.94a.648.648 0 0 0 .002.849l3.868.005v8.024c0 .553.439 1 .982 1h1.965a.99.99 0 0 0 .982-1v-8.02l3.811.005a.65.65 0 0 0-.004-.848L9.414 1.159a.506.506 0 0 0-.764-.001z"/>`),
		g.Group(children),
	)
}
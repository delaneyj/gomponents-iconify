package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.646 10.854a.498.498 0 0 0 .707 0l3.5-3.5a.5.5 0 0 0-.707-.708L12 9.793L8.853 6.647a.5.5 0 0 0-.707.707l3.5 3.5zm3.5 1.792L12 15.793l-3.147-3.146a.5.5 0 0 0-.707.707l3.5 3.5a.498.498 0 0 0 .707 0l3.5-3.5a.5.5 0 0 0-.707-.708z"/>`),
		g.Group(children),
	)
}
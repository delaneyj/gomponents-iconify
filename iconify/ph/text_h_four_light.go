package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHFourLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M254 176a6 6 0 0 1-6 6h-10v26a6 6 0 0 1-12 0v-26h-42a6 6 0 0 1-5.69-7.9l24-72a6 6 0 1 1 11.38 3.8L192.32 170H226v-26a6 6 0 0 1 12 0v26h10a6 6 0 0 1 6 6ZM144 50a6 6 0 0 0-6 6v54H46V56a6 6 0 0 0-12 0v120a6 6 0 0 0 12 0v-54h92v54a6 6 0 0 0 12 0V56a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}
package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17v8a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2v-8m-11 3V3.5M22 9l-6-6l-6 6"/>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-9 7a2 2 0 0 1-2-2v-4l-1-1l1-1V7a2 2 0 0 1 2-2m12 14a2 2 0 0 0 2-2v-4l1-1l-1-1V7a2 2 0 0 0-2-2"/>`),
		g.Group(children),
	)
}
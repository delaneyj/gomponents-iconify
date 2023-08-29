package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZteTaskManager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.057 29.548a2.031 2.031 0 1 1-4.063 0a2.031 2.031 0 0 1 4.063 0ZM9.943 27.875h5.736l1.858-5.02h.758l2.465 6.776h.746l3.827-13.21h.946l4.293 13.21h3.304"/>`),
		g.Group(children),
	)
}
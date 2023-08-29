package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserHandsBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3.905 15.892A4.124 4.124 0 0 1 8 12.25h8a4.124 4.124 0 0 1 4.096 3.642l.649 5.52a.75.75 0 1 1-1.49.176l-.65-5.52A2.624 2.624 0 0 0 16 13.75H8c-1.33 0-2.45.996-2.606 2.317l-.65 5.52a.75.75 0 0 1-1.489-.175l.65-5.52Z" clip-rule="evenodd"/><circle cx="12" cy="6" r="4"/><path d="M8 13.75V18c0 1.886 0 2.828.586 3.414C9.172 22 10.114 22 12 22c1.886 0 2.828 0 3.414-.586C16 20.828 16 19.886 16 18v-4.25H8Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}
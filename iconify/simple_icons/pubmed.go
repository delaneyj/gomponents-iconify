package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pubmed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.23 7.982l.006-1.005C7.846 1.417 5.096 0 5.096 0l.048 2.291C3.73 1.056 2.6 1.444 2.6 1.444l.118 15.307s4.218-1.796 5.428 5.505C10.238 13.535 21.401 24 21.401 24V9S10.52-.18 8.231 7.982zm9.79 9.941l-1.046-5.232l-1.904 4.507h-.96l-1.72-4.301l-1.046 5.04H9.321l2.093-9.39h.802l2.491 5.543l2.508-5.557h.869l2.075 9.39h-2.138z"/>`),
		g.Group(children),
	)
}
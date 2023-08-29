package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M107 384V82l-77 76l-30-30L128 0l128 128l-30 30l-77-76v302h-42z"/>`),
		g.Group(children),
	)
}
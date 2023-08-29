package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Histogram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M4 44h40"/><path stroke-linejoin="round" d="M7 44s5.313-34 17-34c11.688 0 17 34 17 34M4 4v40"/></g>`),
		g.Group(children),
	)
}
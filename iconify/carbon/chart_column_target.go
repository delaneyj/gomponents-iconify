package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartColumnTarget(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 28V6h-8v22h-4V14H7v14H4V2H2v26a2 2 0 0 0 2 2h26v-2zm-14 0H9V16h4zm12 0h-4V8h4zM19 2h8v2h-8z"/><path fill="currentColor" d="M7 10h8v2H7z"/>`),
		g.Group(children),
	)
}
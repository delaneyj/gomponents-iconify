package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeHorizontalCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 10h-1V2h-2v8h-1a2.002 2.002 0 0 0-2 2v8a2.002 2.002 0 0 0 2 2h1v8h2v-8h1a2.002 2.002 0 0 0 2-2v-8a2.002 2.002 0 0 0-2-2zm0 10h-4v-8h4zM12 6h-1V2H9v4H8a2.002 2.002 0 0 0-2 2v16a2.002 2.002 0 0 0 2 2h1v4h2v-4h1a2.002 2.002 0 0 0 2-2V8a2.002 2.002 0 0 0-2-2zm0 18H8V8h4z"/>`),
		g.Group(children),
	)
}
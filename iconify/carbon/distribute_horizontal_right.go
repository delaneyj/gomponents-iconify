package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeHorizontalRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 2h2v28h-2zm-4 20h-4a2.002 2.002 0 0 1-2-2v-8a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v8a2.002 2.002 0 0 1-2 2zm-4-10h-.002L20 20h4v-8zM12 2h2v28h-2zM8 26H4a2.002 2.002 0 0 1-2-2V8a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v16a2.002 2.002 0 0 1-2 2zM4 8h-.002L4 24h4V8z"/>`),
		g.Group(children),
	)
}
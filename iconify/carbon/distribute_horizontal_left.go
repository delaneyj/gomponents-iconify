package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeHorizontalLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 22h-4a2.002 2.002 0 0 1-2-2v-8a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v8a2.002 2.002 0 0 1-2 2zm-4-10v8h4v-8zM18 2h2v28h-2zm-6 24H8a2.002 2.002 0 0 1-2-2V8a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v16a2.002 2.002 0 0 1-2 2zM8 8v16h4V8zM2 2h2v28H2z"/>`),
		g.Group(children),
	)
}
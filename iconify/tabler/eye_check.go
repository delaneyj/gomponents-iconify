package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 12a2 2 0 1 0 4 0a2 2 0 0 0-4 0"/><path d="M11.102 17.957C7.898 17.65 5.198 15.663 3 12c2.4-4 5.4-6 9-6c3.6 0 6.6 2 9 6a19.5 19.5 0 0 1-.663 1.032M15 19l2 2l4-4"/></g>`),
		g.Group(children),
	)
}
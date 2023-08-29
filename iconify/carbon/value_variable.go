package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ValueVariable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 28h-4v-2h4V6h-4V4h4a2.002 2.002 0 0 1 2 2v20a2.002 2.002 0 0 1-2 2zm-6-17h-2l-2 3.897L14 11h-2l2.905 5L12 21h2l2-3.799L18 21h2l-2.902-5L20 11zM10 28H6a2.002 2.002 0 0 1-2-2V6a2.002 2.002 0 0 1 2-2h4v2H6v20h4z"/>`),
		g.Group(children),
	)
}
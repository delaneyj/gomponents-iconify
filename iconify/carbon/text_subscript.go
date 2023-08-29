package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSubscript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 25h-6v-6h4v-2h-4v-2h6v6h-4v2h4v2zM5 7v2h7v16h2V9h7V7H5z"/>`),
		g.Group(children),
	)
}
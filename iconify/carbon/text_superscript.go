package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSuperscript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29 17h-6v-6h4V9h-4V7h6v6h-4v2h4v2zM4 7v2h7v16h2V9h7V7H4z"/>`),
		g.Group(children),
	)
}
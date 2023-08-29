package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toolbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 5c-1.094 0-2 .906-2 2v2H2v18h28V9h-9V7c0-1.094-.906-2-2-2zm0 2h6v2h-6zm-9 4h24v5h-4v-1h-4v1h-8v-1H8v1H4zm0 7h4v1h4v-1h8v1h4v-1h4v7H4z"/>`),
		g.Group(children),
	)
}
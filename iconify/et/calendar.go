package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M34 30.5v-19a.5.5 0 0 0-1 0v19a.5.5 0 0 1-.5.5h-31a.5.5 0 0 1-.5-.5v-19a.5.5 0 0 0-1 0v19c0 .827.673 1.5 1.5 1.5h31c.827 0 1.5-.673 1.5-1.5z"/><path d="M4.5 23a.5.5 0 0 0 0 1H6v3.5a.5.5 0 0 0 1 0V24h6v3.5a.5.5 0 0 0 1 0V24h6v3.5a.5.5 0 0 0 1 0V24h6v3.5a.5.5 0 0 0 1 0V24h1.5a.5.5 0 0 0 0-1H28v-5h1.5a.5.5 0 0 0 0-1H28v-4.5a.5.5 0 0 0-1 0V17h-6v-4.5a.5.5 0 0 0-1 0V17h-6v-4.5a.5.5 0 0 0-1 0V17H7v-4.5a.5.5 0 0 0-1 0V17H4.5a.5.5 0 0 0 0 1H6v5H4.5zM27 18v5h-6v-5h6zm-7 0v5h-6v-5h6zM7 18h6v5H7v-5zM28.5 3h4.25c.087 0 .25 0 .25.5V8H1V3.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 0 0-1h-4C.673 2 0 2.673 0 3.5v5a.5.5 0 0 0 .5.5h33a.5.5 0 0 0 .5-.5v-5c0-1.106-.646-1.5-1.25-1.5H28.5a.5.5 0 0 0 0 1z"/><path d="M22.5 3a.5.5 0 0 0 0-1h-11a.5.5 0 0 0 0 1h11zM9 4.5v-4a.5.5 0 0 0-1 0v4a.5.5 0 0 0 1 0zm16.5.5a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5z"/></g>`),
		g.Group(children),
	)
}
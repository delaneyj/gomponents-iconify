package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Road(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 13h6v2h-6z"/><path fill="currentColor" d="m25.44 8l-1.27-4.55A2.009 2.009 0 0 0 22.246 2H9.754a2.009 2.009 0 0 0-1.923 1.45L6.531 8H4v2h2v7a2.002 2.002 0 0 0 2 2v3h2v-3h12v3h2v-3a2.002 2.002 0 0 0 2-2v-7h2V8zM9.755 4h12.492l1.428 5H8.326zM24 13h-2v2h2v2H8v-2h2v-2H8v-2h16zM2 16h2v14H2zm26 0h2v14h-2z"/><path fill="currentColor" d="M15 21h2v3h-2zm0 5h2v4h-2z"/>`),
		g.Group(children),
	)
}
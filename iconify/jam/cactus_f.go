package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CactusF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 16H0V8.5a2.5 2.5 0 0 1 5 0V11h1V4a4 4 0 1 1 8 0v5h1V6.5a2.5 2.5 0 1 1 5 0V14h-6v6H6v-4zm4-7a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0v-6a1 1 0 0 0-1-1zm0-5a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0V5a1 1 0 0 0-1-1z"/>`),
		g.Group(children),
	)
}
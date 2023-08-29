package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuffleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaShuffleOutline0"><g id="evaShuffleOutline1"><g id="evaShuffleOutline2" fill="currentColor"><path d="M18 9.31a1 1 0 0 0 1 1a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1h-4.3a1 1 0 0 0-1 1a1 1 0 0 0 1 1h1.89L12 10.59L6.16 4.76a1 1 0 0 0-1.41 1.41L10.58 12l-6.29 6.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L18 7.42Z"/><path d="M19 13.68a1 1 0 0 0-1 1v1.91l-2.78-2.79a1 1 0 0 0-1.42 1.42L16.57 18h-1.88a1 1 0 0 0 0 2H19a1 1 0 0 0 1-1.11v-4.21a1 1 0 0 0-1-1Z"/></g></g></g>`),
		g.Group(children),
	)
}
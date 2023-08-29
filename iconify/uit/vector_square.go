package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 16.05v-8.1A2.998 2.998 0 0 0 22 5a2.994 2.994 0 0 0-5.95-.5h-8.1A2.994 2.994 0 0 0 2 5a2.994 2.994 0 0 0 2.5 2.95v8.1A2.994 2.994 0 0 0 5 22a2.998 2.998 0 0 0 2.95-2.5h8.1A2.994 2.994 0 0 0 22 19a2.994 2.994 0 0 0-2.5-2.95zM19 3a2 2 0 1 1 0 4a2 2 0 0 1 0-4zM3 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0zm2 16a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm11.05-2.5h-8.1a2.99 2.99 0 0 0-2.45-2.45v-8.1A2.993 2.993 0 0 0 7.95 5.5h8.1a2.99 2.99 0 0 0 2.45 2.45v8.1a2.99 2.99 0 0 0-2.45 2.45zM19 21a2 2 0 1 1 0-4a2 2 0 0 1 0 4z"/>`),
		g.Group(children),
	)
}
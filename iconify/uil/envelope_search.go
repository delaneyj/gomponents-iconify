package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.21 10.29l-1.73-1.72a4.37 4.37 0 0 0 .65-2.26a4.31 4.31 0 1 0-4.32 4.32a4.37 4.37 0 0 0 2.26-.63l1.72 1.73a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.44ZM18.45 8a2.37 2.37 0 0 1-3.27 0a2.3 2.3 0 0 1-.68-1.64A2.32 2.32 0 0 1 16.81 4a2.3 2.3 0 0 1 1.64.68a2.28 2.28 0 0 1 .68 1.63A2.3 2.3 0 0 1 18.45 8Zm2.05 6a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1V9.41l5.88 5.89a3 3 0 0 0 4.24 0L15 13.88a1 1 0 0 0-1.42-1.42l-1.38 1.42a1 1 0 0 1-1.4 0L4.91 8H9.5a1 1 0 0 0 0-2h-5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}
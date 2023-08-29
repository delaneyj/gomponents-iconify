package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dumbbell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.48 6.55l-2.84-2.84a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41l2.12 2.12l-8.1 8.1l-2.12-2.12a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.42l2.81 2.81l2.81 2.81a1 1 0 0 0 .71.3a1 1 0 0 0 .71-1.71l-2.09-2.09l8.1-8.1l2.12 2.12a1 1 0 1 0 1.41-1.42ZM3.71 17.46a1 1 0 0 0-1.42 1.42l2.83 2.83a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42Zm18-12.34l-2.83-2.83a1 1 0 0 0-1.42 1.42l2.83 2.83a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		g.Group(children),
	)
}
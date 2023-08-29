package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockPattern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 3a4 4 0 0 1 4 4c0 1.86-1.27 3.43-3 3.87v2.26c.37.09.72.24 1.04.43l4.52-4.52C13.2 8.44 13 7.75 13 7a4 4 0 0 1 4-4a4 4 0 0 1 4 4a4 4 0 0 1-4 4c-.74 0-1.43-.2-2-.55L10.45 15c.35.57.55 1.26.55 2a4 4 0 0 1-4 4a4 4 0 0 1-4-4c0-1.86 1.27-3.43 3-3.87v-2.26C4.27 10.43 3 8.86 3 7a4 4 0 0 1 4-4m10 10a4 4 0 0 1 4 4a4 4 0 0 1-4 4a4 4 0 0 1-4-4a4 4 0 0 1 4-4m0 2a2 2 0 0 0-2 2a2 2 0 0 0 2 2a2 2 0 0 0 2-2a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}
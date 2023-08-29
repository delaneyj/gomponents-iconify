package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GolfCourse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4 23.5h5m-2.5 0V0m0 .5h1.879C9.737.5 11.039 1.04 12 2c.96.96 2.263 1.5 3.621 1.5h.879v7h-1.879c-1.358 0-2.66-.54-3.621-1.5a5.121 5.121 0 0 0-3.621-1.5H6.5m17 11s-3.5-4-11.5-4c-1.28 0-2.446.102-3.5.275M.5 18.5s1.261-1.441 4-2.594m11 5.594a2 2 0 1 0 4 0a2 2 0 0 0-4 0Z"/>`),
		g.Group(children),
	)
}
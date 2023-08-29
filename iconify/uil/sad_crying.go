package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SadCrying(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 0 0-3.41 19.39h.06a9.81 9.81 0 0 0 6.7 0h.06A10 10 0 0 0 12 2Zm2 17.74a7.82 7.82 0 0 1-4 0v-3.13a3.79 3.79 0 0 1 4 0Zm2-.82V11a1 1 0 0 0 0-2h-2a1 1 0 0 0 0 2v3.39a6 6 0 0 0-4 0V11a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2v7.92a8 8 0 1 1 8 0Z"/>`),
		g.Group(children),
	)
}
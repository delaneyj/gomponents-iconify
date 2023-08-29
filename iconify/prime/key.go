package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 14.75a5.74 5.74 0 0 1-4.07-1.68A5.77 5.77 0 1 1 15 14.75Zm0-10A4.25 4.25 0 0 0 12 12a4.25 4.25 0 1 0 3-7.25Z"/><path fill="currentColor" d="M4.5 20.25A.74.74 0 0 1 4 20a.75.75 0 0 1 0-1l6.46-6.47a.75.75 0 1 1 1.06 1.07L5 20a.74.74 0 0 1-.5.25Z"/><path fill="currentColor" d="M8 20.75a.74.74 0 0 1-.53-.22l-2-2a.75.75 0 0 1 1.06-1.06l2 2a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22Zm2-2a.74.74 0 0 1-.53-.22l-2-2a.75.75 0 1 1 1.06-1.06l2 2a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22Z"/>`),
		g.Group(children),
	)
}
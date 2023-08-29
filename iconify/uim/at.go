package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func At(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16.5a4.5 4.5 0 1 1 4.5-4.5a4.505 4.505 0 0 1-4.5 4.5Zm0-7a2.5 2.5 0 1 0 2.5 2.5A2.503 2.503 0 0 0 12 9.5Z"/><path fill="currentColor" d="M12 22a10 10 0 1 1 10-10v.75a3.75 3.75 0 0 1-7.5 0V8.5a1 1 0 0 1 2 0v4.25a1.75 1.75 0 0 0 3.5 0V12a8 8 0 1 0-4 6.928a1 1 0 1 1 1 1.733A10.02 10.02 0 0 1 12 22Z"/>`),
		g.Group(children),
	)
}
package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20.75H6a2.64 2.64 0 0 1-2.75-2.53V5.78A2.64 2.64 0 0 1 6 3.25h3a.75.75 0 0 1 0 1.5H6a1.16 1.16 0 0 0-1.25 1v12.47a1.16 1.16 0 0 0 1.25 1h3a.75.75 0 0 1 0 1.5Zm7-4a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06L18.94 12l-3.47-3.47a.75.75 0 1 1 1.06-1.06l4 4a.75.75 0 0 1 0 1.06l-4 4a.74.74 0 0 1-.53.22Z"/><path fill="currentColor" d="M20 12.75H9a.75.75 0 0 1 0-1.5h11a.75.75 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}
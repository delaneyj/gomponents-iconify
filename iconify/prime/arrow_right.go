package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 18.75a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06L17.94 12l-5.47-5.47a.75.75 0 0 1 1.06-1.06l6 6a.75.75 0 0 1 0 1.06l-6 6a.74.74 0 0 1-.53.22Z"/><path fill="currentColor" d="M19 12.75H5a.75.75 0 0 1 0-1.5h14a.75.75 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}
package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 11.75a.74.74 0 0 1-.53-.22L12 6.06l-5.47 5.47a.75.75 0 0 1-1.06-1.06l6-6a.75.75 0 0 1 1.06 0l6 6a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22Z"/><path fill="currentColor" d="M12 19.75a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v14a.76.76 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}
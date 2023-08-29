package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19.75a.74.74 0 0 1-.53-.22l-6-6a.75.75 0 0 1 1.06-1.06L12 17.94l5.47-5.47a.75.75 0 0 1 1.06 1.06l-6 6a.74.74 0 0 1-.53.22Z"/><path fill="currentColor" d="M12 19.75a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v14a.76.76 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}
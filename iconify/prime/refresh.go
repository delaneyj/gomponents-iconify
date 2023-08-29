package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20.75a7.25 7.25 0 0 1 0-14.5h2.5a.75.75 0 0 1 0 1.5H12a5.75 5.75 0 1 0 5.75 5.75a.75.75 0 0 1 1.5 0A7.26 7.26 0 0 1 12 20.75Z"/><path fill="currentColor" d="M12 10.75a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06L13.94 7l-2.47-2.47a.75.75 0 1 1 1.06-1.06l3 3a.75.75 0 0 1 0 1.06l-3 3a.74.74 0 0 1-.53.22Z"/>`),
		g.Group(children),
	)
}
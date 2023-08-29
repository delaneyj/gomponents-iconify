package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.875 8.5h-4.5a1 1 0 0 1 0-2h3.5V3a1 1 0 0 1 2 0v4.5a1 1 0 0 1-1 1Z"/><path fill="currentColor" d="M12 22a10 10 0 1 1 8.651-15.001a1 1 0 0 1-1.73 1.002A7.99 7.99 0 1 0 20 12a1 1 0 0 1 2 0a10.011 10.011 0 0 1-10 10Z"/>`),
		g.Group(children),
	)
}
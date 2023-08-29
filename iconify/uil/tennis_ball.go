package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TennisBall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 0 0-7.35 16.76l.09.1a10 10 0 0 0 14.52 0l.09-.1A10 10 0 0 0 12 2ZM5.61 16.79a7.93 7.93 0 0 1 0-9.58a6 6 0 0 1 0 9.58ZM12 20a7.91 7.91 0 0 1-5-1.77A8 8 0 0 0 7 5.77a7.95 7.95 0 0 1 10 0a8 8 0 0 0 0 12.46A7.91 7.91 0 0 1 12 20Zm6.39-3.21a6 6 0 0 1 0-9.58a7.93 7.93 0 0 1 0 9.58Z"/>`),
		g.Group(children),
	)
}
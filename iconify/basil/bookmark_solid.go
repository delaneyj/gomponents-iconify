package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.13 3.186a25.028 25.028 0 0 0-8.26 0A2.444 2.444 0 0 0 5.877 5.11a36.89 36.89 0 0 0-.148 13.795l.334 1.86a.732.732 0 0 0 1.224.4l4.023-3.822a1 1 0 0 1 1.378 0l4.023 3.822a.732.732 0 0 0 1.224-.4l.334-1.86a36.89 36.89 0 0 0-.148-13.795a2.444 2.444 0 0 0-1.991-1.925Z"/>`),
		g.Group(children),
	)
}
package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.3 5.71a4.92 4.92 0 0 0-3.51-1.46a4.92 4.92 0 0 0-3.51 1.46L12 6l-.28-.28a4.95 4.95 0 0 0-7 0a5 5 0 0 0 0 7l6.77 6.79a.75.75 0 0 0 1.06 0l6.77-6.79a5 5 0 0 0-.02-7.01Z"/>`),
		g.Group(children),
	)
}
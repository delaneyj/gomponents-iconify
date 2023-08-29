package mi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Underline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 4a1 1 0 1 0-2 0v7a4 4 0 0 1-8 0V4a1 1 0 1 0-2 0v7a6 6 0 0 0 12 0V4zM7 19a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2H7z"/>`),
		g.Group(children),
	)
}
package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordMailLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.257 15a5.5 5.5 0 1 1 4.243 2h-13a5.5 5.5 0 1 1 4.243-2h4.514ZM5.5 15a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm13 0a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnArrowRightTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v3.75H5.75a1.75 1.75 0 1 0 0 3.5H7V15a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4Zm13-2a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h1a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2h-1ZM8.97 6.97a.75.75 0 0 1 1.06 0l2 2a.75.75 0 0 1 0 1.06l-2 2a.75.75 0 1 1-1.06-1.06l.72-.72H5.5a.75.75 0 0 1 0-1.5h4.19l-.72-.72a.75.75 0 0 1 0-1.06Z"/>`),
		g.Group(children),
	)
}
package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Overline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.246 18H8.754l-1.6 4H5l6-15h2l6 15h-2.154l-1.6-4Zm-.8-2L12 9.885L9.554 16h4.892ZM4 3h16v2H4V3Z"/>`),
		g.Group(children),
	)
}
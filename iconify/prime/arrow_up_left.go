package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.3 15.54a.75.75 0 0 0 1.5 0V8.86l8.62 8.62a.75.75 0 1 0 1.06-1.06L8.86 7.8h6.68a.75.75 0 0 0 0-1.5H7.05a.75.75 0 0 0-.29.06a.76.76 0 0 0-.46.69Z"/>`),
		g.Group(children),
	)
}
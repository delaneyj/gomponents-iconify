package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.46 6.3a.75.75 0 1 0 0 1.5h6.68l-8.62 8.62a.75.75 0 1 0 1.06 1.06l8.62-8.62v6.68a.75.75 0 0 0 1.5 0V7.05a.75.75 0 0 0-.06-.29a.76.76 0 0 0-.64-.46Z"/>`),
		g.Group(children),
	)
}
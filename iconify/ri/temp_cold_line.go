package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TempColdLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 5a4 4 0 1 1 8 0v5.255a7 7 0 1 1-8 0V5Zm1.144 6.895a5 5 0 1 0 5.712 0L14 11.298V5a2 2 0 1 0-4 0v6.298l-.856.597ZM8 16h8a4 4 0 0 1-8 0Z"/>`),
		g.Group(children),
	)
}
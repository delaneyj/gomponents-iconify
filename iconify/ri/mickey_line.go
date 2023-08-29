package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MickeyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.617 10.913A4.501 4.501 0 0 1 5.5 2a4.5 4.5 0 0 1 4.493 4.254A8.014 8.014 0 0 1 12 6c.693 0 1.365.088 2.007.254a4.5 4.5 0 1 1 5.376 4.66a8 8 0 1 1-14.766 0ZM3 6.5a2.5 2.5 0 0 0 2.766 2.486a8.04 8.04 0 0 1 2.158-1.871A2.5 2.5 0 1 0 3 6.5Zm15.234 2.486a2.5 2.5 0 1 0-2.158-1.871a8.039 8.039 0 0 1 2.158 1.871ZM6 14a6 6 0 1 0 12 0a6 6 0 0 0-12 0Z"/>`),
		g.Group(children),
	)
}
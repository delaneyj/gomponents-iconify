package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3.5a8.5 8.5 0 1 0 8.46 9.32a.5.5 0 0 0-.812-.435a5 5 0 1 1-6.137-7.893a.5.5 0 0 0-.225-.895A8.563 8.563 0 0 0 12 3.5Z"/>`),
		g.Group(children),
	)
}
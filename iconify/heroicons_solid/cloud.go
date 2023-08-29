package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 16a3.5 3.5 0 0 1-.369-6.98a4 4 0 1 1 7.753-1.977A4.5 4.5 0 1 1 13.5 16h-8Z"/>`),
		g.Group(children),
	)
}
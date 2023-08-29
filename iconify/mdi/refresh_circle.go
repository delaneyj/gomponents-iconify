package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefreshCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m6 9h-5l1.81-1.81A3.94 3.94 0 0 0 12 8a4 4 0 1 0 3.86 5h2.05A6 6 0 1 1 12 6a5.91 5.91 0 0 1 4.22 1.78L18 6Z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YinYang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 0 1 10 10a10 10 0 0 1-10 10A10 10 0 0 1 2 12A10 10 0 0 1 12 2m0 2a8 8 0 0 0-8 8a8 8 0 0 0 8 8a4 4 0 0 1-4-4a4 4 0 0 1 4-4a4 4 0 0 0 4-4a4 4 0 0 0-4-4m0 2.5A1.5 1.5 0 0 1 13.5 8A1.5 1.5 0 0 1 12 9.5A1.5 1.5 0 0 1 10.5 8A1.5 1.5 0 0 1 12 6.5m0 8a1.5 1.5 0 0 0-1.5 1.5a1.5 1.5 0 0 0 1.5 1.5a1.5 1.5 0 0 0 1.5-1.5a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		g.Group(children),
	)
}
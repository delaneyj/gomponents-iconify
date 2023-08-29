package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2h-1v20h1a10 10 0 0 0 0-20zm1 17.94V4.06a8 8 0 0 1 0 15.88z"/>`),
		g.Group(children),
	)
}
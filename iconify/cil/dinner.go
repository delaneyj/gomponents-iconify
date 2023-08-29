package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dinner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M464 344.063c0-109.308-84.755-199.193-192-207.39V80h-32v56.673c-107.245 8.2-192 98.082-192 207.39v33.107h416Zm-32 1.107H80v-1.107c0-97.046 78.953-176 176-176s176 78.953 176 176ZM16 416h480v32H16z"/>`),
		g.Group(children),
	)
}
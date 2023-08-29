package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineRotateNinetyDegreesCw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 13a9 9 0 0 0 13.79 7.62l-1.46-1.46c-.99.53-2.13.84-3.33.84c-3.86 0-7-3.14-7-7s3.14-7 7-7h.17L9.59 7.59L11 9l4-4l-4-4l-1.42 1.41L11.17 4H11a9 9 0 0 0-9 9zm9 0l6 6l6-6l-6-6l-6 6zm6 3.17L13.83 13L17 9.83L20.17 13L17 16.17z"/>`),
		g.Group(children),
	)
}
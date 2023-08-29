package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangularFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#F8312F" d="M27.626 11.944L8 4H7v17h1l19.626-7.944a.6.6 0 0 0 0-1.112Z"/><path fill="#E39D89" d="M4 4a2 2 0 1 1 4 0v26H4V4Z"/></g>`),
		g.Group(children),
	)
}
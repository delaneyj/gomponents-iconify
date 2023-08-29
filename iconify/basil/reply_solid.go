package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.446 16.06a.5.5 0 0 1-.655.68l-2.5-1.153a14.381 14.381 0 0 0-6.681-1.309a61.43 61.43 0 0 1-.121 2.204l-.069.938a.754.754 0 0 1-1.158.581a19.55 19.55 0 0 1-5.351-5.068l-.46-.64a.5.5 0 0 1 0-.584l.46-.64A19.55 19.55 0 0 1 9.262 6a.754.754 0 0 1 1.158.58l.069.94c.046.63.082 1.26.108 1.89h.644a9.5 9.5 0 0 1 8.475 5.209l.73 1.442Z"/>`),
		g.Group(children),
	)
}
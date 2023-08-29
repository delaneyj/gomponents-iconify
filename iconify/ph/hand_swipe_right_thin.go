package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandSwipeRightThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 140v36c0 24.66-8.08 41.1-8.42 41.79a4 4 0 1 1-7.16-3.58c.07-.15 7.58-15.55 7.58-38.21v-36a16 16 0 0 0-32 0v4a4 4 0 0 1-8 0v-20a16 16 0 0 0-32 0v12a4 4 0 0 1-8 0V68a16 16 0 0 0-32 0v108a4 4 0 0 1-7.39 2.11l-18.68-30a.75.75 0 0 1-.07-.12a16 16 0 0 0-27.72 16l29.31 50a4 4 0 0 1-6.9 4L23.22 168a24 24 0 0 1 41.52-24.09L76 162V68a24 24 0 0 1 48 0v38.13a24 24 0 0 1 39.94 16.06A24 24 0 0 1 204 140Zm38.83-86.83l-32-32a4 4 0 0 0-5.66 5.66L230.34 52H168a4 4 0 0 0 0 8h62.34l-25.17 25.17a4 4 0 0 0 5.66 5.66l32-32a4 4 0 0 0 0-5.66Z"/>`),
		g.Group(children),
	)
}
package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Httpie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.28 0C4.4 0 1.992 2.279 1.933 5.155a5.263 5.263 0 0 0 5.26 5.358h4.223a.306.306 0 0 1 .122.584l-6.47 2.835a5.263 5.263 0 0 0-3.135 4.85C1.953 21.683 4.368 24 7.273 24h2.212c2.922 0 5.357-2.345 5.35-5.267a5.263 5.263 0 0 0-3.29-4.867a.303.303 0 0 1-.007-.556l7.402-3.246a5.263 5.263 0 0 0 3.128-4.846C22.047 2.317 19.626.003 16.724.003z"/>`),
		g.Group(children),
	)
}
package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Directions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.25a1.81 1.81 0 0 1-1.28-.53l-7.44-7.44a1.81 1.81 0 0 1 0-2.56l7.44-7.44a1.86 1.86 0 0 1 2.56 0l7.44 7.44a1.81 1.81 0 0 1 0 2.56l-7.44 7.44a1.81 1.81 0 0 1-1.28.53Zm0-17a.27.27 0 0 0-.21.09l-7.45 7.45a.29.29 0 0 0 0 .42l7.45 7.45a.25.25 0 0 0 .42 0l7.45-7.45a.29.29 0 0 0 0-.42l-7.45-7.45a.27.27 0 0 0-.21-.09Z"/><path fill="currentColor" d="M13.27 14.42a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l1.7-1.68l-1.7-1.68a.75.75 0 1 1 1.06-1.06l2.2 2.21A.75.75 0 0 1 16 12l-2.2 2.2a.74.74 0 0 1-.53.22Z"/><path fill="currentColor" d="M8.5 15a.76.76 0 0 1-.75-.75v-2.79a.76.76 0 0 1 .75-.75h7a.75.75 0 0 1 0 1.5H9.25v2a.76.76 0 0 1-.75.79Z"/>`),
		g.Group(children),
	)
}
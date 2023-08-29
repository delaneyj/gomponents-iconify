package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapMarkerEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.46 9.63A8.5 8.5 0 1 0 6 16.46l5.3 5.31a1 1 0 0 0 1.42 0L18 16.46a8.46 8.46 0 0 0 2.46-6.83Zm-3.86 5.42l-4.6 4.6l-4.6-4.6a6.49 6.49 0 0 1-1.87-5.22A6.57 6.57 0 0 1 8.42 5a6.47 6.47 0 0 1 7.16 0a6.57 6.57 0 0 1 2.89 4.81a6.49 6.49 0 0 1-1.87 5.24Zm-2.81-8.8a1 1 0 0 0-1.42 0L8.79 9.83a1 1 0 0 0-.29.7V13a1 1 0 0 0 1 1h2.42a1 1 0 0 0 .71-.29l3.58-3.58a1 1 0 0 0 0-1.41ZM11.51 12h-1v-1l2.58-2.58l1 1Z"/>`),
		g.Group(children),
	)
}
package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundedrectangleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M11.5 4.5c-6.939 0-11 4.47-11 11v11c0 6.971 3.859 11 11 11h18c7.4 0 11-4.029 11-11v-11c0-6.97-4.061-11-11-11h-18zm-1.01 3H30.5c4.59 0 7 2.52 7 7v13c0 4.62-2.45 7-7 7h-20c-4.54 0-7-2.49-7-7V14.37c0-4.38 2.529-6.87 6.99-6.87z"/>`),
		g.Group(children),
	)
}
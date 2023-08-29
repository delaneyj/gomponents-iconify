package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 3.63a3.57 3.57 0 0 1 2.49-1.06a2.9 2.9 0 0 1 3 3A3.65 3.65 0 0 1 10.73 9l1 1a4.71 4.71 0 0 0 1.54-1a4.79 4.79 0 0 0 1.42-3.38a4.16 4.16 0 0 0-1.22-3a4.16 4.16 0 0 0-3-1.22a4.79 4.79 0 0 0-3.38 1.34a4.71 4.71 0 0 0-1 1.54l1 1A3.6 3.6 0 0 1 8 3.63zm-6.31-.17l2.59 2.59a4.71 4.71 0 0 0-1.54 1a4.79 4.79 0 0 0-1.42 3.38a4.16 4.16 0 0 0 1.22 3a4.16 4.16 0 0 0 3 1.22a4.79 4.79 0 0 0 3.38-1.42a4.71 4.71 0 0 0 1-1.54l2.74 2.74l.89-.88l-11-11zM8 12.37a3.57 3.57 0 0 1-2.49 1.06a2.9 2.9 0 0 1-3-3A3.65 3.65 0 0 1 5.26 7L9 10.74a3.57 3.57 0 0 1-1 1.63z"/>`),
		g.Group(children),
	)
}
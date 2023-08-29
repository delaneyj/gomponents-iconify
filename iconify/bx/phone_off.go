package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.09 12.5a8.92 8.92 0 0 1-1-2.2l1.59-1.59a1 1 0 0 0 0-1.42l-4-4a1 1 0 0 0-1.41 0L2.59 6A2 2 0 0 0 2 7.44A15.44 15.44 0 0 0 5.62 17L2.3 20.29l1.41 1.42l18-18l-1.41-1.42zM7 15.55a13.36 13.36 0 0 1-3-8.13l2-2L8.59 8L7.3 9.29a1 1 0 0 0-.27.92a11 11 0 0 0 1.62 3.73zm9.71-2.26a1 1 0 0 0-1.41 0l-1.6 1.6l-.34-.12l-1.56 1.55a12.06 12.06 0 0 0 2 .66a1 1 0 0 0 .91-.27l1.3-1.3L18.59 18l-2 2A13.61 13.61 0 0 1 10 18.1l-1.43 1.45a15.63 15.63 0 0 0 8 2.45a2 2 0 0 0 1.43-.58l2.71-2.71a1 1 0 0 0 0-1.42z"/>`),
		g.Group(children),
	)
}
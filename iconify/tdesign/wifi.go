package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.193 9.808c-5.077-5.077-13.308-5.077-18.385 0l-.707.707L.687 9.1l.707-.707c5.857-5.857 15.355-5.857 21.213 0l.707.707l-1.414 1.415l-.707-.707Zm-4.597 4.596a6.5 6.5 0 0 0-9.192 0l-.707.707l-1.414-1.414l.707-.707a8.5 8.5 0 0 1 12.02 0l.708.707l-1.415 1.414l-.707-.707Zm-6.01 3.182a2 2 0 0 1 2.829 0l.707.707L12 20.414l-2.12-2.121l.707-.707Z"/>`),
		g.Group(children),
	)
}
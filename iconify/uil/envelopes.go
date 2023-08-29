package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envelopes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 21.63H5a3 3 0 0 1-3-3v-8a1 1 0 0 0-2 0v8a5 5 0 0 0 5 5h12a1 1 0 0 0 0-2Zm4-18H7a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3Zm-.41 2l-5.88 5.88a1 1 0 0 1-1.42 0L7.41 5.63Zm1.41 11a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7l5.88 5.88a3 3 0 0 0 4.24 0L22 7Z"/>`),
		g.Group(children),
	)
}
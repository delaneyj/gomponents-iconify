package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicAccidentalDoubleSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.41 10H17V7h-3v1.59l-2 2l-2-2V7H7v3h1.59l2 2l-2 2H7v3h3v-1.59l2-2l2 2V17h3v-3h-1.59l-2-2l2-2Z"/>`),
		g.Group(children),
	)
}
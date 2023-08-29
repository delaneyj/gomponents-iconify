package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicAccidentalSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 9.5v-2l-2 .6V5.5h-2v3.2l-2 .6V6.5H9v3.4l-2 .6v2l2-.6v2l-2 .6v2l2-.6v2.6h2v-3.2l2-.6v2.8h2v-3.4l2-.6v-2l-2 .6v-2l2-.6m-4 3.2l-2 .6v-2l2-.6v2Z"/>`),
		g.Group(children),
	)
}
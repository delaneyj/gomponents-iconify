package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitlabAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 21.42l3.684-11.333H8.32L12 21.421Z"/><path fill="currentColor" d="m3.16 10.087l-1.123 3.444a.763.763 0 0 0 .277.852l9.685 7.038l-8.84-11.334Z" opacity=".25"/><path fill="currentColor" d="M3.16 10.087h5.16L6.1 3.262a.383.383 0 0 0-.728 0L3.16 10.087Z"/><path fill="currentColor" d="m20.845 10.087l1.118 3.444a.763.763 0 0 1-.276.852l-9.688 7.038l8.846-11.334Z" opacity=".25"/><path fill="currentColor" d="M20.845 10.087h-5.161L17.9 3.262a.383.383 0 0 1 .727 0l2.217 6.825Z"/><path fill="currentColor" d="m11.999 21.421l3.685-11.334h5.161l-8.846 11.334z" opacity=".5"/><path fill="currentColor" d="m11.999 21.421l-8.84-11.334H8.32l3.679 11.334z" opacity=".5"/>`),
		g.Group(children),
	)
}
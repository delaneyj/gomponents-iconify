package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Baffled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zM8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13z"/><path fill="currentColor" d="M6 6.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0z"/><path fill="currentColor" d="M5.5 5a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 5.5 5zm0-1C4.122 4 3 5.122 3 6.5S4.122 9 5.5 9S8 7.878 8 6.5S6.878 4 5.5 4zM11 6.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0z"/><path fill="currentColor" d="M10.5 5a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 10.5 5zm0-1C9.121 4 8 5.122 8 6.5S9.121 9 10.5 9S13 7.878 13 6.5S11.879 4 10.5 4zM6 11h4v1H6v-1z"/>`),
		g.Group(children),
	)
}
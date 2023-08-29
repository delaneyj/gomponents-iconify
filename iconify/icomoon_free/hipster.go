package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hipster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zM8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13zM4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0zm6 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0z"/><path fill="currentColor" d="M10.561 8.439a1.5 1.5 0 1 0-2.063 2.176c1.352 1.227 4.503-.029 4.503-1.615c-.969.625-1.726.153-2.439-.561z"/><path fill="currentColor" d="M5.439 8.439a1.5 1.5 0 1 1 2.063 2.176C6.15 11.842 2.999 10.586 2.999 9c.969.625 1.726.153 2.439-.561z"/>`),
		g.Group(children),
	)
}
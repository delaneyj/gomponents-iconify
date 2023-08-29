package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ello(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zm4.885 9.212C12.31 11.442 10.301 13 8 13s-4.31-1.558-4.885-3.788a.708.708 0 0 1 .685-.884c.322 0 .604.218.684.531a3.631 3.631 0 0 0 7.032 0a.707.707 0 1 1 1.369.354z"/>`),
		g.Group(children),
	)
}
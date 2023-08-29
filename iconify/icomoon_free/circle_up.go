package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 8a8 8 0 1 0 16 0A8 8 0 0 0 0 8zm14.5 0a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0z"/><path fill="currentColor" d="m11.043 10.457l1.414-1.414L8 4.586L3.543 9.043l1.414 1.414L8 7.414z"/>`),
		g.Group(children),
	)
}
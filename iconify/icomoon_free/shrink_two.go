package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShrinkTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7 9v6.5L4.5 13l-3 3L0 14.5l3-3L.5 9zm9-7.5l-3 3L15.5 7H9V.5L11.5 3l3-3z"/>`),
		g.Group(children),
	)
}
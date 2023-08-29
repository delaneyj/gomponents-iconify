package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAmountAsc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5 12V0H3v12H.5L4 15.5L7.5 12H5z"/><path fill="currentColor" d="M7 9h9v2H7V9zm0-3h7v2H7V6zm0-3h5v2H7V3zm0-3h3v2H7V0z"/>`),
		g.Group(children),
	)
}
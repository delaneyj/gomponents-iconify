package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Creditcard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 3v18H1V3h22Zm-2 2H3v4h18V5Zm0 6H3v8h18v-8Zm-11 5H5v-2h5v2Z"/>`),
		g.Group(children),
	)
}
package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adidas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m1.33 19l-.6-1.036l4.33-2.5L7.103 19H1.329Zm13.856 0H9.412l-3.619-6.268l4.33-2.5L15.187 19Zm8.083 0h-5.774l-6.64-11.5l4.33-2.5l8.084 14Z"/>`),
		g.Group(children),
	)
}
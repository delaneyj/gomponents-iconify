package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.54.889l-3.63 1.773v18.17c-4.15-.52-7.27-2.78-7.27-5.5c0-2.58 2.8-4.75 6.63-5.41v-2.31C4.42 8.322 0 11.502 0 15.332c0 3.96 4.74 7.24 10.91 7.78l3.63-1.71V.888m.64 6.724v2.31c1.43.25 2.71.7 3.76 1.31l-1.97 1.11l7.03 1.53l-.5-5.21l-1.87 1.06c-1.74-1.06-3.96-1.81-6.45-2.11z"/>`),
		g.Group(children),
	)
}
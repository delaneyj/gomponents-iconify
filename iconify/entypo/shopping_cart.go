package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 17a2 2 0 1 0 3.999.001A2 2 0 0 0 13 17zM3 17a2 2 0 1 0 4 0a2 2 0 0 0-4 0zm3.547-4.828L17.615 9.01A.564.564 0 0 0 18 8.5V3H4V1.4c0-.22-.181-.4-.399-.4H.399A.401.401 0 0 0 0 1.4V3h2l1.91 8.957l.09.943v1.649c0 .219.18.4.4.4h13.2c.22 0 .4-.182.4-.4V13H6.752c-1.15 0-1.174-.551-.205-.828z"/>`),
		g.Group(children),
	)
}
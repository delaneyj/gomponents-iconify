package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOffSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12.254 12.961l1.892 1.893a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708L4.45 5.156c-.215.414-.36.87-.419 1.351A3.25 3.25 0 0 0 4.25 13h7.5c.171 0 .34-.013.504-.039ZM15 9.75c0 .867-.34 1.654-.892 2.237l-8.32-8.32a4 4 0 0 1 6.182 2.84A3.25 3.25 0 0 1 15 9.75Z"/>`),
		g.Group(children),
	)
}
package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aids(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.47 13.321l-2.339 3.774l4.281 6.904l3.37-2.086zM13.555 0H2.225L-.001 3.97l5.313 8.58l2.339-3.781l-2.972-4.8h6.418L-.001 21.913l3.376 2.086L15.78 3.969z"/>`),
		g.Group(children),
	)
}
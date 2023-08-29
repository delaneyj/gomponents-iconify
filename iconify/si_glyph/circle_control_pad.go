package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleControlPad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 .032a8 8 0 1 0 0 16a8 8 0 0 0 0-16zM5 9.859a.755.755 0 0 1-.88 0L2.182 8.345c-.243-.19-.243-.498 0-.688L4.12 6.143a.75.75 0 0 1 .88 0v3.716zm1.142-5.74l1.514-1.937c.19-.243.498-.243.687 0L9.856 4.12a.744.744 0 0 1 0 .881H6.143a.754.754 0 0 1-.001-.881zm3.715 7.742L8.344 13.8c-.193.243-.5.243-.689 0l-1.512-1.939a.748.748 0 0 1 0-.879h3.714a.746.746 0 0 1 0 .879zm2.024-2.004a.751.751 0 0 1-.881 0V6.143a.751.751 0 0 1 .88 0l1.937 1.513c.244.191.244.5 0 .689L11.88 9.857z"/>`),
		g.Group(children),
	)
}
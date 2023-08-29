package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coupon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21H3V1.132l1.555 1.036l2.391 1.594l1.93-1.543l.624-.5l.625.5L12 3.72l1.875-1.5l.625-.5l.625.5l1.929 1.543l2.391-1.594L21 1.132V21H4ZM5 4.87V19h14V4.87l-1.445.963l-.609.406l-.57-.457l-1.876-1.5l-1.875 1.5l-.625.5l-.625-.5l-1.875-1.5l-1.875 1.5l-.571.457l-.609-.406L5 4.87ZM8 16H7v-2h10v2H8Zm3-4h6v-2h-6v2Zm-3-2H7v2.004h2.004V10H8Z"/>`),
		g.Group(children),
	)
}
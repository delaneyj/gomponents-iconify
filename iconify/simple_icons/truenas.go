package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Truenas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 10.049v5.114l-10.949 6.324v-5.114L24 10.049zm-24 0v5.114l10.956 6.324v-5.114L0 10.049zm12.004-.605l-4.433 2.559l4.433 2.559l4.429-2.559l-4.429-2.559zm10.952-1.207l-9.905-5.723v5.118l5.473 3.164l4.432-2.559zm-12-.605V2.513L1.044 8.236l4.432 2.555l5.48-3.159z"/>`),
		g.Group(children),
	)
}
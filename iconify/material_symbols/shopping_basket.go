package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBasket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.525 21q-.675 0-1.2-.413T3.6 19.525l-2.55-9.25Q.925 9.8 1.213 9.4T2 9h4.75l4.4-6.55q.125-.2.35-.325T11.975 2q.25 0 .475.125t.35.325L17.2 9H22q.5 0 .788.4t.162.875l-2.55 9.25q-.2.65-.725 1.063t-1.2.412H5.525ZM12 17q.825 0 1.413-.588T14 15q0-.825-.588-1.413T12 13q-.825 0-1.413.588T10 15q0 .825.588 1.413T12 17ZM9.175 9H14.8l-2.825-4.2l-2.8 4.2Z"/>`),
		g.Group(children),
	)
}
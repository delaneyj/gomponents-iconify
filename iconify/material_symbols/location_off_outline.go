package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 10q0-.825-.588-1.413T12 8q-.25 0-.475.05t-.425.175l2.675 2.675q.125-.2.175-.425T14 10Zm4.2 5.325l-1.45-1.45q.625-1.05.938-1.963T18 10.2q0-2.725-1.738-4.462T12 4q-1.1 0-2.063.338T8.2 5.325L6.775 3.9Q7.85 2.975 9.2 2.488T12 2q3.175 0 5.588 2.225T20 10.2q0 1.2-.45 2.463t-1.35 2.662ZM14.275 17.1L6.1 8.925q-.05.3-.075.625T6 10.2q0 1.775 1.475 4.063T12 19.35q.65-.575 1.213-1.137t1.062-1.113Zm6.2 6.2l-4.8-4.8q-.8.85-1.7 1.725T12 22q-4.025-3.425-6.012-6.362T4 10.2q0-.8.125-1.525t.35-1.375l-3.8-3.8L2.1 2.075l19.8 19.8l-1.425 1.425ZM10.2 13.025ZM12.475 9.6Z"/>`),
		g.Group(children),
	)
}
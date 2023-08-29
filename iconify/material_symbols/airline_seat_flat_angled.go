package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatFlatAngled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.025 12.725l2.4-6.575l8.45 3.075q1.575.575 2.275 2.063t.125 3.062l-1.025 2.825l-12.225-4.45Zm-7.6.425l.675-1.875l18.8 6.85L20.225 20l-18.8-6.85Zm4.9-1.6q-1.25 0-2.125-.875T3.325 8.55q0-1.25.875-2.125t2.125-.875q1.25 0 2.125.875t.875 2.125q0 1.25-.875 2.125t-2.125.875Z"/>`),
		g.Group(children),
	)
}
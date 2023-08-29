package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatFlatAngledOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.05 12.725l2.4-6.575l8.45 3.075q1.575.575 2.275 2.063t.125 3.062l-1.025 2.825l-12.225-4.45Zm3.6-4q-.2.5-.363.938t-.337.937l-.325.925l1.025-2.8ZM1.45 13.15l.675-1.875l18.8 6.85L20.25 20l-18.8-6.85Zm4.875-1.6q-1.25 0-2.125-.875T3.325 8.55q0-1.25.875-2.125t2.125-.875q1.25 0 2.125.875t.875 2.125q0 1.25-.875 2.125t-2.125.875Zm0-2q.425 0 .713-.288t.287-.712q0-.425-.288-.713t-.712-.287q-.425 0-.713.288t-.287.712q0 .425.288.713t.712.287Zm5.3 1.975l8.45 3.1l.35-.95q.275-.8-.062-1.55T19.225 11.1L12.65 8.725l-1.025 2.8Zm-5.3-2.95Z"/>`),
		g.Group(children),
	)
}
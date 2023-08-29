package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatFlatAngledOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.275 17.175L9.05 12.725l2.4-6.575l12.225 4.45l-2.4 6.575Zm-8.625-8.45l-1.025 2.8l1.025-2.8ZM20.25 20l-18.8-6.85l.675-1.875l18.8 6.85L20.25 20ZM6.325 11.55q-1.25 0-2.125-.875T3.325 8.55q0-1.25.875-2.125t2.125-.875q1.25 0 2.125.875t.875 2.125q0 1.25-.875 2.125t-2.125.875Zm0-2q.425 0 .712-.288t.288-.712q0-.425-.288-.713t-.712-.287q-.425 0-.713.288t-.287.712q0 .425.288.713t.712.287Zm5.3 1.975l8.45 3.1l1.05-2.85l-8.475-3.05l-1.025 2.8Zm-5.3-2.95Z"/>`),
		g.Group(children),
	)
}
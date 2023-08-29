package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsAltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.375 0-.738-.15t-.662-.45l-8-8q-.3-.3-.45-.663T2 12q0-.375.15-.738t.45-.662l8-8q.3-.3.663-.45T12 2q.375 0 .738.15t.662.45l8 8q.3.3.45.663T22 12q0 .375-.15.738t-.45.662l-8 8q-.3.3-.663.45T12 22Zm1.15-9l-1.875 1.9q-.275.275-.275.688t.3.712q.275.275.7.275t.7-.275l3.6-3.6q.3-.3.3-.7t-.3-.7l-3.6-3.6q-.275-.275-.687-.288T11.3 7.7q-.275.275-.288.688t.263.712L13.15 11H8q-.425 0-.713.288T7 12q0 .425.288.713T8 13h5.15Z"/>`),
		g.Group(children),
	)
}
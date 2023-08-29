package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinchZoomOutOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 12V7h1.5v2.45l2.975-2.975l1.05 1.05L3.55 10.5H6V12H1Zm6.525-5.475l-1.05-1.05L9.45 2.5H7V1h5v5h-1.5V3.55L7.525 6.525ZM14.825 23q-.6 0-1.15-.225t-.975-.65L7.6 17l.75-.775q.4-.4.938-.537t1.062.012l1.65.475V8q0-.425.288-.713T13 7q.425 0 .713.288T14 8v10.825l-2.425-.675l2.55 2.55q.125.125.313.213t.387.087H19q.825 0 1.413-.588T21 19v-4q0-.425.288-.713T22 14q.425 0 .713.288T23 15v4q0 1.65-1.175 2.825T19 23h-4.175ZM15 16v-4q0-.425.288-.713T16 11q.425 0 .713.288T17 12v4h-2Zm3 0v-3q0-.425.288-.713T19 12q.425 0 .713.288T20 13v3h-2Zm-.5 2Z"/>`),
		g.Group(children),
	)
}
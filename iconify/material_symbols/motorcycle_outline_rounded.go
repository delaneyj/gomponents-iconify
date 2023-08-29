package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MotorcycleOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.7 11H8.95h3.75h-2ZM5 19q-2.075 0-3.538-1.463T0 14q0-2.075 1.463-3.538T5 9h11.6l-2-2H12q-.425 0-.713-.288T11 6q0-.425.288-.713T12 5h2.575q.4 0 .763.15t.637.425L19.45 9.05q1.95.15 3.25 1.575T24 14q0 2.075-1.463 3.538T19 19q-2.075 0-3.538-1.463T14 14q0-.45.063-.888t.237-.862l-2.45 2.45q-.15.15-.337.225t-.388.075H9.9q-.35 1.75-1.725 2.875T5 19Zm14-2q1.25 0 2.125-.875T22 14q0-1.25-.875-2.125T19 11q-1.25 0-2.125.875T16 14q0 1.25.875 2.125T19 17ZM5 17q.95 0 1.713-.55T7.8 15H6q-.425 0-.712-.288T5 14q0-.425.288-.713T6 13h1.8q-.325-.9-1.088-1.45T5 11q-1.25 0-2.125.875T2 14q0 1.25.875 2.125T5 17Zm4.95-4h.75l2-2H8.95q.375.425.625.925T9.95 13Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Living(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v16q0 .825-.588 1.413T20 22H4Zm3-4h10q.825 0 1.413-.588T19 16v-3.5q0-.675-.338-1.238t-.912-.912V9q0-1.25-.875-2.125T14.75 6h-5.5Q8 6 7.125 6.875T6.25 9v1.35q-.575.35-.913.913T5 12.5V16q0 .825.588 1.413T7 18Zm0-1.5q-.225 0-.363-.138T6.5 16v-3.5q0-.425.288-.713T7.5 11.5q.425 0 .713.288t.287.712v2h7v-2q0-.425.288-.713t.712-.287q.425 0 .713.288t.287.712V16q0 .225-.138.363T17 16.5H7Zm3-3.5v-.5q0-.975-.65-1.675t-1.6-.775V9q0-.65.425-1.075T9.25 7.5h5.5q.65 0 1.075.425T16.25 9v1.05q-.95.075-1.6.775T14 12.5v.5h-4Z"/>`),
		g.Group(children),
	)
}
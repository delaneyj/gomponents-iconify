package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TouchAppOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.475 22q-.7 0-1.312-.3t-1.038-.85l-5.45-6.925l.475-.5q.5-.525 1.2-.625t1.3.275L7.5 14.2V6q0-.425.288-.713T8.5 5q.425 0 .725.288t.3.712v11.8L7.1 16.3l2.6 3.325q.15.175.35.275t.425.1H16q.825 0 1.413-.588T18 18v-4q0-.425-.288-.713T17 13h-5.475v-2H17q1.25 0 2.125.875T20 14v4q0 1.65-1.175 2.825T16 22h-5.525Zm-6.3-13.5q-.325-.55-.5-1.188T3.5 6q0-2.075 1.463-3.538T8.5 1q2.075 0 3.538 1.463T13.5 6q0 .675-.175 1.313t-.5 1.187l-1.725-1q.2-.35.3-.713T11.5 6q0-1.25-.875-2.125T8.5 3q-1.25 0-2.125.875T5.5 6q0 .425.1.788t.3.712l-1.725 1Zm8.375 7Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedtimeOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.075 8.25L9.025 6.2l.1-.838q.05-.413.15-.862q-.35.125-.675.275t-.65.35L6.5 3.65q.975-.65 1.937-1t2.013-.5q.725-.1 1.05.35t.05 1.25q-.4 1.125-.512 2.25t.037 2.25Zm8 13.65L17.5 20.325q-1.2.8-2.575 1.238T12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-1.525.425-2.913T3.65 6.5L2.1 4.925q-.275-.275-.275-.688t.275-.712q.3-.3.713-.3t.712.3L20.5 20.5q.3.3.288.7t-.313.7q-.3.275-.7.288t-.7-.288ZM12 20q1.1 0 2.125-.288t1.925-.837L5.125 7.95q-.525.9-.825 1.912T4 12q0 3.325 2.337 5.663T12 20Zm-2.85-8.025ZM9.025 6.2Z"/>`),
		g.Group(children),
	)
}
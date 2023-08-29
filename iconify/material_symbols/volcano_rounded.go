package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolcanoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.25 19.2l2.225-5q.25-.55.738-.875T7.3 13H9l1.5-3.75q.225-.575.725-.912T12.35 8h4.15q.675 0 1.2.4t.725 1.05l2.85 10q.275.95-.325 1.75t-1.6.8H5.075q-1.1 0-1.688-.9t-.137-1.9ZM13 4V2q0-.425.288-.713T14 1q.425 0 .713.288T15 2v2q0 .425-.288.713T14 5q-.425 0-.713-.288T13 4Zm5.225 2.175q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l1.45-1.45q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-1.45 1.45Zm-8.45 0l-1.45-1.45q-.275-.275-.263-.7t.288-.7q.275-.275.7-.275t.7.275l1.425 1.45q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275Z"/>`),
		g.Group(children),
	)
}
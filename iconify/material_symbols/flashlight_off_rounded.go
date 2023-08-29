package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 22q-.825 0-1.413-.588T8 20v-9.2L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L16 18.8V20q0 .825-.588 1.413T14 22h-4Zm6-8.85L9.85 7H18v1l-2 3v2.15ZM18 5H7.85L6.125 3.275q.225-.575.725-.925T8 2h8q.825 0 1.413.587T18 4v1Z"/>`),
		g.Group(children),
	)
}
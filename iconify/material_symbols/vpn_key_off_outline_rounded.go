package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VpnKeyOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.075 21.9L12.15 15q-.8 1.35-2.138 2.175T7 18q-2.5 0-4.25-1.75T1 12q0-1.65.8-3.025T3.95 6.8L2.075 4.925q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l17 17q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3ZM17 14.175L15.825 13H17v1.175Zm3.45 3.4l-1.45-1.4V13h2v-2h-7.175l-2-2H21q.825 0 1.413.588T23 11v2q0 .825-.588 1.413T21 15v1.175q0 .45-.163.8t-.387.6ZM7 16q1.625 0 2.487-.875t1.238-1.575L5.475 8.3q-1.1.45-1.787 1.438T3 12q0 1.65 1.175 2.825T7 16Zm0-2q-.825 0-1.413-.587T5 12q0-.825.588-1.413T7 10q.825 0 1.413.588T9 12q0 .825-.588 1.413T7 14Zm7.725-2.1Zm-7.85.25Z"/>`),
		g.Group(children),
	)
}
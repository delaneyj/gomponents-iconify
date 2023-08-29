package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterMediumOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.375 14.3L7 20h10l.65-6h-.225q-.95 0-1.738-.138t-2.137-.537q-.575-.175-1.2-.25T11.1 13q-1.275 0-2.462.325t-2.263.975Zm-.25-2.15q1.175-.575 2.438-.863T11.124 11q.75 0 1.488.1t1.462.3q1.25.35 1.913.475T17.4 12h.475l.875-8H5.25l.875 8.15Zm.85 9.85q-.775 0-1.338-.5T5 20.225L3 2h18l-2 18.225q-.075.775-.638 1.275t-1.337.5H6.975ZM7 20h10H7Z"/>`),
		g.Group(children),
	)
}
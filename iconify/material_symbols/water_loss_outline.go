package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterLossOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.1 16q-1.175 0-2.287.288t-2.138.837L7 20h10l.325-3q-.9 0-1.675-.15t-2.1-.525q-.575-.175-1.2-.25T11.1 16Zm-4.65-1q1.125-.5 2.3-.75t2.375-.25q.75 0 1.488.1t1.462.3q1.25.35 1.913.475T17.4 15h.15l1.2-11H5.25l1.2 11Zm.525 7q-.775 0-1.337-.5T5 20.225L3 2h18l-2 18.225q-.075.775-.638 1.275t-1.337.5H6.975Zm4.125-2H17H7h4.1Z"/>`),
		g.Group(children),
	)
}
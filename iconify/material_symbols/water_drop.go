package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterDrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.275 19q.3-.025.513-.238T13 18.25q0-.35-.225-.563T12.2 17.5q-1.025.075-2.175-.563t-1.45-2.312q-.05-.275-.262-.45T7.825 14q-.35 0-.575.263t-.15.612q.425 2.275 2 3.25t3.175.875ZM12 22q-3.425 0-5.713-2.35T4 13.8q0-2.5 1.988-5.438T12 2q4.025 3.425 6.013 6.363T20 13.8q0 3.5-2.288 5.85T12 22Z"/>`),
		g.Group(children),
	)
}
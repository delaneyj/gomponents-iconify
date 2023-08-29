package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodPressureSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 22q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22Zm.2-4.5l2.275-2.275l-.7-.7L16.5 16.8l.7.7ZM2 11V4h20v8.1q-1-1.025-2.288-1.563T17 10q-.95 0-1.813.238t-1.637.662L11.9 7.55q-.125-.25-.375-.375T11 7.05q-.275 0-.525.125t-.375.375L7 13.775L5.875 11.55q-.125-.25-.363-.4T5 11H2Zm0 9v-7h2.375L6.1 16.45q.125.275.363.413T7 17q.275 0 .513-.138t.362-.412L11 10.25l.95 1.9q-.9.95-1.425 2.188T10 17q0 .775.163 1.538T10.675 20H2Z"/>`),
		g.Group(children),
	)
}
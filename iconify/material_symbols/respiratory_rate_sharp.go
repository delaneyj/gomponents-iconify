package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RespiratoryRateSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 23q-.625 0-1.063-.438T11 21.5v-3.675l1.325-3.525q.225-.575.738-.938T14.2 13h.8v2.55l-1.35 1.35l.7.7L16 15.925V11h2v4.925l1.65 1.675l.7-.7L19 15.525V13h.8q.625 0 1.137.363t.738.937L23 17.825V21.5q0 .625-.438 1.063T21.5 23h-2q-.625 0-1.063-.438T18 21.5v-2.75l-1-1l-1 1v2.75q0 .625-.438 1.063T14.5 23h-2ZM7 13.775L5.875 11.55q-.125-.25-.363-.4T5 11H2V4h20v7.5h-2.5v-2h-5v2l-.625.025L11.9 7.55q-.125-.25-.375-.375T11 7.05q-.275 0-.525.125t-.375.375L7 13.775ZM2 20v-7h2.375L6.1 16.45q.125.275.363.413T7 17q.275 0 .513-.138t.362-.412L11 10.25l1 2.025q-.35.3-.638.675t-.462.825L9.5 17.55V20H2Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VpnLockSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.05 0-3.875-.788t-3.188-2.15q-1.362-1.362-2.15-3.187T2 12q0-2.075.788-3.888t2.15-3.174Q6.3 3.575 8.124 2.788T12 2q.8 0 1.538.113T15 2.45V7h-4v3H8v2h7v4h1l1.9 1.4q.975-1.1 1.538-2.463T20 12q0-.275-.025-.5T19.9 11h2.05q.05.275.05.5v.5q0 2.05-.787 3.875t-2.15 3.188q-1.363 1.362-3.175 2.15T12 22Zm-1-2.05V18l-2-2v-1l-4.8-4.8q-.075.45-.138.9T4 12q0 3.1 2.013 5.338T11 19.95ZM17 9V4h1V3q0-.825.588-1.413T20 1q.825 0 1.413.588T22 3v1h1v5h-6Zm2-5h2V3q0-.425-.288-.713T20 2q-.425 0-.713.288T19 3v1Z"/>`),
		g.Group(children),
	)
}
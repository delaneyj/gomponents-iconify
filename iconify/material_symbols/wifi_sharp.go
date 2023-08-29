package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.35 15.35l-2.1-2.15q1.55-1.55 3.55-2.375T12 10q2.2 0 4.213.838t3.537 2.412l-2.1 2.1q-1.125-1.125-2.588-1.738T12 13q-1.6 0-3.063.613T6.35 15.35ZM2.1 11.1L0 9q2.375-2.425 5.488-3.713T12 4q3.4 0 6.513 1.288T24 9l-2.1 2.1q-1.975-1.975-4.538-3.038T12 7Q9.2 7 6.637 8.063T2.1 11.1ZM12 21l-3.525-3.55q.7-.7 1.613-1.075T12 16q1 0 1.913.375t1.612 1.075L12 21Z"/>`),
		g.Group(children),
	)
}
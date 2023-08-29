package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 14q-1.25 0-2.125-.875T9 11V5q0-1.25.875-2.125T12 2q1.25 0 2.125.875T15 5v6q0 1.25-.875 2.125T12 14Zm0 7q-.425 0-.713-.288T11 20v-2.1q-2.325-.3-3.95-1.925t-1.975-3.9q-.075-.425.225-.75T6.1 11q.35 0 .625.262t.35.638q.325 1.75 1.7 2.925T12 16q1.85 0 3.225-1.175t1.7-2.925q.075-.375.362-.638t.638-.262q.475 0 .775.325t.225.75q-.35 2.275-1.975 3.9T13 17.9V20q0 .425-.288.713T12 21Z"/>`),
		g.Group(children),
	)
}
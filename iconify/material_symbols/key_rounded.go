package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 15q1.25 0 2.125-.875T10 12q0-1.25-.875-2.125T7 9q-1.25 0-2.125.875T4 12q0 1.25.875 2.125T7 15Zm0 3q-2.5 0-4.25-1.75T1 12q0-2.5 1.75-4.25T7 6q2.025 0 3.538 1.15T12.65 10h7.95q.2 0 .388.075t.312.2l.975.975q.15.15.225.337t.075.388q0 .2-.062.375t-.213.325l-2.6 2.6q-.15.15-.325.225t-.375.075q-.2 0-.375-.063T18.3 15.3L17 14l-1.3 1.3q-.15.15-.325.212t-.375.063q-.2 0-.375-.063T14.3 15.3L13 14h-.35q-.625 1.8-2.175 2.9T7 18Z"/>`),
		g.Group(children),
	)
}
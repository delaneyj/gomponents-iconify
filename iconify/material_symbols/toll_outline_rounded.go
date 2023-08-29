package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TollOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 20q-3.35 0-5.675-2.325T7 12q0-3.35 2.325-5.675T15 4q3.35 0 5.675 2.325T23 12q0 3.35-2.325 5.675T15 20Zm-9.375-.75q-2.1-.975-3.363-2.925T1 12q0-2.375 1.263-4.325T5.625 4.75q.525-.25.95.038T7 5.724q0 .25-.163.488t-.412.362q-1.575.725-2.5 2.188T3 12q0 1.775.925 3.238t2.5 2.187q.25.125.413.35t.162.5q0 .625-.425.925t-.95.05ZM15 12Zm0 6q2.5 0 4.25-1.75T21 12q0-2.5-1.75-4.25T15 6q-2.5 0-4.25 1.75T9 12q0 2.5 1.75 4.25T15 18Z"/>`),
		g.Group(children),
	)
}
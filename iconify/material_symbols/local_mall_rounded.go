package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalMallRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20V8q0-.825.588-1.413T5 6h2q0-2.075 1.463-3.538T12 1q2.075 0 3.538 1.463T17 6h2q.825 0 1.413.588T21 8v12q0 .825-.588 1.413T19 22H5ZM9 6h6q0-1.25-.875-2.125T12 3q-1.25 0-2.125.875T9 6Zm2.975 8q1.85 0 3.35-1.225t1.475-2.75q0-.425-.275-.725t-.7-.3q-.35 0-.625.225t-.4.675q-.275.95-1.075 1.525t-1.75.575q-.95 0-1.762-.575T9.15 9.9q-.125-.475-.375-.688T8.175 9q-.425 0-.713.3t-.287.725q0 1.525 1.475 2.75T11.975 14Z"/>`),
		g.Group(children),
	)
}
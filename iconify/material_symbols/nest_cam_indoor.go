package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestCamIndoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 11q-.825 0-1.413-.588T10 9q0-.825.588-1.413T12 7q.825 0 1.413.588T14 9q0 .825-.588 1.413T12 11ZM7 22q0-.825.588-1.413T9 20h2q0-1.575-.763-2.925T8.15 14.85q-1.5-.975-2.325-2.525T5 9q0-2.925 2.038-4.963T12 2q2.925 0 4.963 2.038T19 9q0 1.8-.85 3.35t-2.35 2.525q-1.325.85-2.063 2.2T13 20h2q.825 0 1.413.588T17 22H7Zm5-8q2.075 0 3.538-1.463T17 9q0-2.075-1.463-3.538T12 4Q9.925 4 8.462 5.463T7 9q0 2.075 1.463 3.538T12 14Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22v-5.925q0-.5.238-.95T4.9 14.4l1.1-.725q.175 2.1.55 3.575t1.175 3.275L4 22Zm5.225-2q-.875-1.65-1.3-3.5T7.5 12.675q0-3.125 1.238-5.888T12 2.6q2.025 1.425 3.263 4.188t1.237 5.887q0 1.95-.425 3.788T14.775 20h-5.55ZM12 13q.825 0 1.413-.588T14 11q0-.825-.588-1.413T12 9q-.825 0-1.413.588T10 11q0 .825.588 1.413T12 13Zm8 9l-3.725-1.475q.8-1.8 1.175-3.275t.55-3.575l1.1.725q.425.275.663.725t.237.95V22Z"/>`),
		g.Group(children),
	)
}
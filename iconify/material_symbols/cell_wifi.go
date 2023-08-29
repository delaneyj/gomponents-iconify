package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellWifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22L22 5.95V22H6Zm12-2h2v-9.2l-2 2V20ZM7.85 9.8l-1.3-1.3q.95-.95 2.1-1.4t2.4-.45q1.25 0 2.4.45t2.1 1.4l-1.3 1.3q-.675-.675-1.5-1.012t-1.7-.338q-.875 0-1.7.338T7.85 9.8ZM5.3 7.2L4 5.95q1.475-1.475 3.3-2.212T11.05 3q1.925 0 3.775.738T18.15 5.95l-1.3 1.25q-1.2-1.2-2.713-1.775T11.05 4.85q-1.575 0-3.063.575T5.3 7.2Zm5.75 5.8l-1.925-1.925q.4-.4.9-.613t1.025-.212q.525 0 1.012.2t.913.625L11.05 13Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpaRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 20.85q.05.475-.313.813t-.812.262q-4.35-.65-6.9-3.363t-2.9-7.387q-.05-.5.288-.838t.837-.262q4.575.625 7.05 3.525T13 20.85ZM11.225 2.975q.3-.425.788-.412T12.8 3q1.125 1.575 1.963 3.45t1.012 3.15q-.975.45-2.175 1.463t-1.575 1.562q-.35-.55-1.613-1.613T8.275 9.6q.2-1.25 1.012-3.162t1.938-3.463Zm9.65 7.1q.475-.05.8.263t.3.812q-.2 4.025-2.188 6.538T14.976 21.2q-.05-1.525-.462-3.462T13.225 14.4q1.075-1.65 3.188-2.85t4.462-1.475Z"/>`),
		g.Group(children),
	)
}
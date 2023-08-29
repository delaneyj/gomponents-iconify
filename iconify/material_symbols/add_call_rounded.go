package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddCallRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.95 21q-3.125 0-6.175-1.363t-5.55-3.862q-2.5-2.5-3.862-5.55T3 4.05q0-.45.3-.75t.75-.3H8.1q.35 0 .625.238t.325.562l.65 3.5q.05.4-.025.675T9.4 8.45L6.975 10.9q.5.925 1.187 1.787t1.513 1.663q.775.775 1.625 1.438T13.1 17l2.35-2.35q.225-.225.588-.338t.712-.062l3.45.7q.35.1.575.363T21 15.9v4.05q0 .45-.3.75t-.75.3ZM16 8h-2q-.425 0-.713-.288T13 7q0-.425.288-.713T14 6h2V4q0-.425.288-.713T17 3q.425 0 .713.288T18 4v2h2q.425 0 .713.288T21 7q0 .425-.288.713T20 8h-2v2q0 .425-.288.713T17 11q-.425 0-.713-.288T16 10V8Z"/>`),
		g.Group(children),
	)
}
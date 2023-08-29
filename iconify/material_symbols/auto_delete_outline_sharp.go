package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoDeleteOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 6v13V6Zm4.25 15H5V6H4V4h5V3h6v1h5v2h-1v4.3q-.425-.125-.988-.213T17 10V6H7v13h3.3q.15.525.4 1.038t.55.962ZM9 17h1q0-1.575.5-2.588L11 13.4V8H9v9Zm4-5.75q.425-.275.963-.55T15 10.3V8h-2v3.25ZM17 22q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22Zm1.65-2.65l.7-.7l-1.85-1.85V14h-1v3.2l2.15 2.15Z"/>`),
		g.Group(children),
	)
}
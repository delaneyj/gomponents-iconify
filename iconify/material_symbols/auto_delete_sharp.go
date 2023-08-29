package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoDeleteSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21V6H4V4h5V3h6v1h5v2h-1v4.3q-.425-.125-.988-.213T17 10q-.45 0-1.012.075t-.988.2V8h-2v3.25q-.55.4-1.1.987T11 13.4V8H9v9h1q0 .975.35 2.087t.9 1.913H5Zm12 1q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22Zm1.65-2.65l.7-.7l-1.85-1.85V14h-1v3.2l2.15 2.15Z"/>`),
		g.Group(children),
	)
}
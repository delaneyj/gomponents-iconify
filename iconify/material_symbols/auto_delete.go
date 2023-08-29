package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 3h6v1h5v2h-1v4.3q-.425-.125-.988-.213T17 10q-.45 0-1.012.075t-.988.2V8h-2v3.25q-.55.4-1.1.987T11 13.4V8H9v9h1q0 .975.35 2.087t.9 1.913H7q-.825 0-1.413-.588T5 19V6H4V4h5V3Zm8 9q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12Zm-.5 2v3.2l2.15 2.15l.7-.7l-1.85-1.85V14h-1Z"/>`),
		g.Group(children),
	)
}
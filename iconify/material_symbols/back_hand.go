package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackHand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.8 23q-2.05 0-3.85-.938T6 19.45L1.2 12.4l.475-.475q.5-.525 1.238-.6t1.337.35l2.75 1.9V4q0-.425.288-.713T8 3q.425 0 .713.288T9 4v8h2V2q0-.425.288-.713T12 1q.425 0 .713.288T13 2v10h2V3q0-.425.288-.713T16 2q.425 0 .713.288T17 3v9h2V5q0-.425.288-.713T20 4q.425 0 .713.288T21 5v9.8q0 3.425-2.388 5.813T12.8 23Z"/>`),
		g.Group(children),
	)
}
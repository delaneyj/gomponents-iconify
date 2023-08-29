package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestTag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.925 0-4.963-2.038T5 15q0-2.35 1.375-4.2T10 8.3V3q0-.425.288-.713T11 2h2q.425 0 .713.288T14 3v5.3q2.225.65 3.613 2.5T19 15q0 2.925-2.05 4.963T12 22Zm0-2q2.075 0 3.538-1.45T17 15q0-2.075-1.463-3.538T12 10q-2.1 0-3.55 1.463T7 15q0 2.1 1.45 3.55T12 20Z"/>`),
		g.Group(children),
	)
}
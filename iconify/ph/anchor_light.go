package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 138a6 6 0 0 0-6 6a42 42 0 0 1-42 42a45.86 45.86 0 0 0-34 15.08V126h34a6 6 0 0 0 0-12h-34V85.4a30 30 0 1 0-12 0V114H88a6 6 0 0 0 0 12h34v75.08A45.86 45.86 0 0 0 88 186a42 42 0 0 1-42-42a6 6 0 0 0-12 0a54.06 54.06 0 0 0 54 54a34 34 0 0 1 34 34a6 6 0 0 0 12 0a34 34 0 0 1 34-34a54.06 54.06 0 0 0 54-54a6 6 0 0 0-6-6ZM110 56a18 18 0 1 1 18 18a18 18 0 0 1-18-18Z"/>`),
		g.Group(children),
	)
}
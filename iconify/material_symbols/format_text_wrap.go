package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatTextWrap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20V4h2v16H4Zm14 0V4h2v16h-2Zm-7.4-2.45L7.05 14l3.55-3.525l1.4 1.4L10.875 13H13q.825 0 1.413-.588T15 11q0-.825-.588-1.413T13 9H7V7h6q1.65 0 2.825 1.175T17 11q0 1.65-1.175 2.825T13 15h-2.125L12 16.125l-1.4 1.425Z"/>`),
		g.Group(children),
	)
}
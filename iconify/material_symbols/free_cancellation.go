package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreeCancellation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.55 22.5L13 18.95l1.4-1.4l2.125 2.125l4.25-4.25l1.4 1.425l-5.625 5.65ZM7.4 17L6 15.6L7.6 14L6 12.4L7.4 11L9 12.6l1.6-1.6l1.4 1.4l-1.6 1.6l1.6 1.6l-1.4 1.4L9 15.4L7.4 17ZM5 22q-.825 0-1.413-.588T3 20V6q0-.825.588-1.413T5 4h1V2h2v2h8V2h2v2h1q.825 0 1.413.588T21 6v6.35l-2 2.025V10H5v10h6.25l1.975 2H5Z"/>`),
		g.Group(children),
	)
}
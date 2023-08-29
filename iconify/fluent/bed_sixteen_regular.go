package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 4.5A2.5 2.5 0 0 1 5.5 2h5A2.5 2.5 0 0 1 13 4.5v1.55a2.5 2.5 0 0 1 2 2.45v5a.5.5 0 0 1-1 0V11H2v2.5a.5.5 0 0 1-1 0v-5a2.5 2.5 0 0 1 2-2.45V4.5ZM2 10h12V8.5A1.5 1.5 0 0 0 12.5 7h-9A1.5 1.5 0 0 0 2 8.5V10Zm10-5.5A1.5 1.5 0 0 0 10.5 3h-5A1.5 1.5 0 0 0 4 4.5V6h1v-.5a.5.5 0 0 1 .5-.5H7a.5.5 0 0 1 .5.5V6h1v-.5A.5.5 0 0 1 9 5h1.5a.5.5 0 0 1 .5.5V6h1V4.5Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenRotationUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.75 23.25l-1.4-1.425L14.175 20H8q-.825 0-1.413-.588T6 18V7.825l2 2V18h6.175l-1.825-1.825l1.4-1.425L18 19l-4.25 4.25ZM18 16.175l-2-2V6H9.825l1.825 1.825l-1.4 1.425L6 5L10.25.75l1.4 1.425L9.825 4H16q.825 0 1.413.588T18 6v10.175Z"/>`),
		g.Group(children),
	)
}
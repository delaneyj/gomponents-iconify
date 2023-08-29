package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicineBottleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 2v2h-2v3a3 3 0 0 1 3 3v11a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V10a3 3 0 0 1 3-3V4H5V2h14Zm-2 7H7a1 1 0 0 0-1 1v10h12V10a1 1 0 0 0-1-1Zm-4 2v2h2v2h-2.001L13 17h-2l-.001-2H9v-2h2v-2h2Zm2-7H9v3h6V4Z"/>`),
		g.Group(children),
	)
}
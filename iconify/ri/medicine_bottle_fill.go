package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicineBottleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 5v2a3 3 0 0 1 3 3v11a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V10a3 3 0 0 1 3-3V5h10Zm-4 6h-2v2H9v2h1.999L11 17h2l-.001-2H15v-2h-2v-2Zm6-9v2H5V2h14Z"/>`),
		g.Group(children),
	)
}
package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DriveLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.097 6.15L4.31 14.443l1.755 3.031l4.785-8.288L9.097 6.15Zm-1.3 12.324h9.568l1.751-3.034H9.55l-1.752 3.034Zm11.314-5.034l-4.786-8.29H10.83l4.787 8.29h3.495ZM8.52 3.15h6.96L22 14.444l-3.48 6.03H5.49L2 14.444L8.52 3.15Zm3.485 8.036l-1.302 2.254h2.603l-1.301-2.254Z"/>`),
		g.Group(children),
	)
}
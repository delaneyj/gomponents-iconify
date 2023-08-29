package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 4H5v16h14V4ZM3 2.992C3 2.444 3.447 2 3.998 2H20a1 1 0 0 1 1 1v17.992A1 1 0 0 1 20.007 22H3.993A1 1 0 0 1 3 21.008V2.992Zm8.293 10.13l4.243-4.243l1.414 1.414l-5.657 5.657l-3.89-3.89l1.415-1.414l2.475 2.475Z"/>`),
		g.Group(children),
	)
}
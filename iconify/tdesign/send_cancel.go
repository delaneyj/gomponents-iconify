package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendCancel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.292 1.665L24.002 12l-5.456 2.379l-.8-1.833L18.998 12L3.708 5.336l2 5.664H11v2H5.708l-2 5.665l8.56-3.731l.799 1.833L.292 22.336L3.94 12L.292 1.665Zm22.79 14.33l-2.832 2.851l2.821 2.822l-1.414 1.414l-2.819-2.82l-2.818 2.819l-1.414-1.415l2.818-2.817l-2.838-2.838L16 14.597l2.836 2.835l2.827-2.846l1.419 1.41Z"/>`),
		g.Group(children),
	)
}
package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadarFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.368 4.398l-3.484 6.035l1.732 1l3.485-6.035c4.169 2.772 6.305 7.08 4.56 10.102c-1.86 3.222-7.19 3.355-11.91.63C4.028 13.403 1.48 8.721 3.34 5.5c1.745-3.023 6.543-3.327 11.028-1.102Zm1.516-2.625l1.732 1l-1.5 2.598l-1.732-1l1.5-2.598ZM6.732 20H17v2H5.018a.998.998 0 0 1-1.015-.922a.995.995 0 0 1 .131-.578l2.25-3.897l1.732 1L6.732 20Z"/>`),
		g.Group(children),
	)
}
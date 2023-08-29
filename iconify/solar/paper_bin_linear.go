package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperBinLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.034 8.89c-.48-3.204-.721-4.805.176-5.848C4.108 2 5.728 2 8.967 2h6.066c3.24 0 4.86 0 5.757 1.042c.897 1.043.657 2.644.177 5.848l-1.2 8c-.366 2.438-.549 3.656-1.393 4.383c-.844.727-2.076.727-4.541.727h-3.666c-2.465 0-3.697 0-4.541-.727c-.844-.727-1.027-1.945-1.392-4.383l-1.2-8Z"/><path stroke-linejoin="round" d="m8 6l-4.5 5l7.5 8m3-13L4 16M20 6L7 19m6 0l7.5-8L16 6m-6 0l10 10M4 6l13 13"/><path stroke-linecap="round" d="M21 6H3m16 13H5"/></g>`),
		g.Group(children),
	)
}
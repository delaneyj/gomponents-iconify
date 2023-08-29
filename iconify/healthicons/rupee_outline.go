package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RupeeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 12a3 3 0 0 1 3-3h12a3 3 0 1 1 0 6a3 3 0 1 1 0 6h-1.512a9.024 9.024 0 0 1-5.276 5.41l5.284 7.926a3 3 0 0 1-4.992 3.328l-8-12A3 3 0 0 1 18 21a3 3 0 0 1 0-6a3 3 0 0 1-3-3Zm3-1a1 1 0 1 0 0 2h2a4.998 4.998 0 0 1 4.33 2.5a1 1 0 0 1-.865 1.5H18a1 1 0 1 0 0 2h5.465a1 1 0 0 1 .866 1.5A4.998 4.998 0 0 1 20 23h-2a1 1 0 0 0-.832 1.555l8 12a1 1 0 0 0 1.664-1.11l-6.037-9.056a1 1 0 0 1 .63-1.534a7.012 7.012 0 0 0 5.354-5.104a1 1 0 0 1 .969-.751H30a1 1 0 1 0 0-2h-2.252a1 1 0 0 1-.969-.75a6.949 6.949 0 0 0-.715-1.75a1 1 0 0 1 .866-1.5H30a1 1 0 1 0 0-2H18Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
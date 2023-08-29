package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibleIconAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 7a2 2 0 1 0-2-2a2 2 0 0 0 2 2Zm-3.3 11.4A4 4 0 1 1 9.05 12a1 1 0 1 0-.22-2a6 6 0 0 0 .67 12a6 6 0 0 0 4.8-2.4a1 1 0 0 0-1.6-1.2Zm6.8 1.6h-1v-5a1 1 0 0 0-1-1h-4.57l1.69-4.66v-.31a1.1 1.1 0 0 0 0-.18a1.06 1.06 0 0 0 0-.19a1.4 1.4 0 0 0-.09-.17a.72.72 0 0 0-.11-.15a.64.64 0 0 0-.15-.13s0-.06-.08-.08L9.71 5.55h-.12a1.06 1.06 0 0 0-.19-.06H9a.8.8 0 0 0-.2.07h-.11L6 7.13A1 1 0 0 0 6.48 9A1 1 0 0 0 7 8.87l2.23-1.3l3.24 1.88l-1.89 5.21a.88.88 0 0 0 0 .16a.58.58 0 0 0 0 .18a3 3 0 0 0 .08.38l.11.15a.57.57 0 0 0 .11.16a.67.67 0 0 0 .14.09a1.22 1.22 0 0 0 .19.12a1 1 0 0 0 .34.06h5v5a1 1 0 0 0 1 1h2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}
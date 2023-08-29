package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeScanLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M10 22c-3.771 0-5.657 0-6.828-1.172C2 19.657 2 18.771 2 15m20 0c0 3.771 0 4.657-1.172 5.828C19.657 22 17.771 22 14 22m0-20c3.771 0 5.657 0 6.828 1.172C22 4.343 22 5.229 22 9M10 2C6.229 2 4.343 2 3.172 3.172C2 4.343 2 5.229 2 9"/><path d="M5.892 14.06C5.297 13.37 5 13.025 5 12c0-1.025.297-1.37.892-2.06C7.08 8.562 9.072 7 12 7c2.927 0 4.92 1.562 6.108 2.94c.595.69.892 1.035.892 2.06c0 1.025-.297 1.37-.892 2.06C16.92 15.438 14.928 17 12 17c-2.927 0-4.92-1.562-6.108-2.94Z"/><circle cx="12" cy="12" r="2"/></g>`),
		g.Group(children),
	)
}
package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScannerLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M10 22c-3.771 0-5.657 0-6.828-1.172C2 19.657 2 18.771 2 15m20 0c0 3.771 0 4.657-1.172 5.828C19.657 22 17.771 22 14 22m0-20c3.771 0 5.657 0 6.828 1.172C22 4.343 22 5.229 22 9M10 2C6.229 2 4.343 2 3.172 3.172C2 4.343 2 5.229 2 9" opacity=".5"/><path d="M2 12h20"/></g>`),
		g.Group(children),
	)
}
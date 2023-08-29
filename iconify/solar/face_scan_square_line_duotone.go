package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceScanSquareLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M9 16c.85.63 1.885 1 3 1s2.15-.37 3-1"/><ellipse cx="15" cy="10.5" fill="currentColor" rx="1" ry="1.5"/><ellipse cx="9" cy="10.5" fill="currentColor" rx="1" ry="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M22 14c0 3.771 0 5.657-1.172 6.828C19.657 22 17.771 22 14 22m-4 0c-3.771 0-5.657 0-6.828-1.172C2 19.657 2 17.771 2 14m8-12C6.229 2 4.343 2 3.172 3.172C2 4.343 2 6.229 2 10m12-8c3.771 0 5.657 0 6.828 1.172C22 4.343 22 6.229 22 10" opacity=".5"/></g>`),
		g.Group(children),
	)
}
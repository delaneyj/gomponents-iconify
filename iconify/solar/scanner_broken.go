package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScannerBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M10 22H9m-7-7c0 3.771 0 4.657 1.172 5.828c.653.654 1.528.943 2.828 1.07M22 15c0 3.771 0 4.657-1.172 5.828C19.657 22 17.771 22 14 22m0-20h1m7 7c0-3.771 0-4.657-1.172-5.828c-.653-.654-1.528-.943-2.828-1.07M10 2C6.229 2 4.343 2 3.172 3.172C2 4.343 2 5.229 2 9m0 3h6m14 0H12"/>`),
		g.Group(children),
	)
}
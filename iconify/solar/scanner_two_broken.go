package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScannerTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M6 13s1.8-2 6-2s6 2 6 2m-8 9c-3.771 0-5.657 0-6.828-1.172C2 19.657 2 17.771 2 14M14 2c3.771 0 5.657 0 6.828 1.172C22 4.343 22 6.229 22 10m0 4v1m-8 7c3.771 0 5.657 0 6.828-1.172c.654-.653.943-1.528 1.07-2.828M2 10V9m8-7C6.229 2 4.343 2 3.172 3.172C2.518 3.825 2.229 4.7 2.102 6"/>`),
		g.Group(children),
	)
}
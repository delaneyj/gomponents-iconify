package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrawingPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.329 1.136a.5.5 0 0 0-.708.707l.653.653l-4.848 3.637l-1.108-1.108a.5.5 0 0 0-.707.707l1.414 1.414l1.06 1.061l-3.27 3.27a.5.5 0 1 0 .708.708l3.27-3.27l1.06 1.06l1.415 1.414a.5.5 0 1 0 .707-.707L8.867 9.574l3.637-4.848l.653.653a.5.5 0 1 0 .707-.707l-1.06-1.061l-1.415-1.414l-1.06-1.06Zm-4.19 5.711l4.85-3.637l.8.801l-3.636 4.85L6.14 6.846Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
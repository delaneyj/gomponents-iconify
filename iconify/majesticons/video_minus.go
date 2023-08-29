package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v1.586l2.293-2.293A1 1 0 0 1 22 8v8a1 1 0 0 1-1.707.707L18 14.414V16a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V8zm5 3a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H7z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCloseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 2a1 1 0 0 0-1 1v16a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8a1 1 0 0 0-.293-.707l-5-5A1 1 0 0 0 14 2H5Zm9 2.414L17.586 8H14V4.414Zm-3.293 7.879a1 1 0 0 0-1.414 1.414L10.586 15l-1.293 1.293a1 1 0 1 0 1.414 1.414L12 16.414l1.293 1.293a1 1 0 0 0 1.414-1.414L13.414 15l1.293-1.293a1 1 0 0 0-1.414-1.414L12 13.586l-1.293-1.293Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
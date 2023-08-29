package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buysellads(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 9.419l2.681 10.044h-5.362zM30 5v22c0 1.656-1.344 3-3 3H5c-1.656 0-3-1.344-3-3V5c0-1.656 1.344-3 3-3h22c1.656 0 3 1.344 3 3zm-4.081 20.331L20.013 6.662h-8.025L6.082 25.331h5.669l6.981-5.725l1.512 5.725z"/>`),
		g.Group(children),
	)
}
package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeOpenSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 3l-.531.344l-12 7.812L3 11.47V29h26V11.469l-.469-.313l-12-7.812zm0 2.375L26.188 12L16 18.594L5.812 12zM5 13.844l10.469 6.781l.531.344l.531-.344L27 13.844V27H5z"/>`),
		g.Group(children),
	)
}
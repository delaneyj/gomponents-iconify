package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PyramidMaya(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 2h12v2h-1v4h2v3h1v3h1v3h1v6H2v-6h1v-3h1v-3h1V8h2V4H6V2Zm3 2v4h5.999L15 4H9Zm4 6h-2v11h2V10Zm2 11h5v-2h-1v-3h-1v-3h-1v-3h-2v11Zm-6 0V10H7v3H6v3H5v3H4v2h5Zm2-16.002h2.004v2.004H11V4.998Z"/>`),
		g.Group(children),
	)
}
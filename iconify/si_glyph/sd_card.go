package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M5 9h5.906v4.938H5z"/><path d="M12.969 5.938h.969V1.063H5L2.022 6.215v1.754h1.009v1.053H2.022v7.916h11.915V9.022h-.969V5.938h.001zM10 1.969h1V4h-1V1.969zm-2 0h1V4H8V1.969zM6 2h1.031v2.031H6V2zm6 13.025H3.993V7.937H12v7.088zM13 4h-1V1.969h1V4z"/></g>`),
		g.Group(children),
	)
}
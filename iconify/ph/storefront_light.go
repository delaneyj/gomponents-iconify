package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorefrontLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M230 96a6.19 6.19 0 0 0-.22-1.65l-14.34-50.2A14.07 14.07 0 0 0 202 34H54a14.07 14.07 0 0 0-13.43 10.15l-14.34 50.2A6.19 6.19 0 0 0 26 96v16a38 38 0 0 0 16 31v65a14 14 0 0 0 14 14h144a14 14 0 0 0 14-14v-65a38 38 0 0 0 16-31ZM52.11 47.45A2 2 0 0 1 54 46h148a2 2 0 0 1 1.92 1.45L216.05 90H40ZM102 102h52v10a26 26 0 0 1-52 0Zm-64 0h52v10a26 26 0 0 1-52 0Zm164 106a2 2 0 0 1-2 2H56a2 2 0 0 1-2-2v-59.34a38 38 0 0 0 42-16.21a37.95 37.95 0 0 0 64 0a38 38 0 0 0 42 16.21Zm-10-70a26 26 0 0 1-26-26v-10h52v10a26 26 0 0 1-26 26Z"/>`),
		g.Group(children),
	)
}
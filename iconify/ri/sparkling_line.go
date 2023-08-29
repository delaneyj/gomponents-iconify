package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SparklingLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 4.438A2.437 2.437 0 0 0 16.438 2h1.125A2.437 2.437 0 0 0 20 4.438v1.125A2.437 2.437 0 0 0 17.562 8h-1.125A2.437 2.437 0 0 0 14 5.562V4.438ZM1 11a6 6 0 0 0 6-6h2a6 6 0 0 0 6 6v2a6 6 0 0 0-6 6H7a6 6 0 0 0-6-6v-2Zm3.876 1A8.038 8.038 0 0 1 8 15.124A8.038 8.038 0 0 1 11.124 12A8.038 8.038 0 0 1 8 8.876A8.038 8.038 0 0 1 4.876 12Zm12.374 2A3.25 3.25 0 0 1 14 17.25v1.5A3.25 3.25 0 0 1 17.25 22h1.5A3.25 3.25 0 0 1 22 18.75v-1.5A3.25 3.25 0 0 1 18.75 14h-1.5Z"/>`),
		g.Group(children),
	)
}
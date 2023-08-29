package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortBoolAscendingVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 17h3l-4 4l-4-4h3V3h2v14M9 13H5c-1.11 0-2 .89-2 2v4c0 1.11.89 2 2 2h4c1.11 0 2-.89 2-2v-4c0-1.11-.89-2-2-2m-2.73 6.5l-2.53-2.55l1.07-1.05l1.47 1.49L9.2 14.5l1.06 1.05l-3.99 3.95M9 3H5c-1.11 0-2 .89-2 2v4c0 1.11.89 2 2 2h4c1.11 0 2-.89 2-2V5c0-1.11-.89-2-2-2m0 6H5V5h4v4Z"/>`),
		g.Group(children),
	)
}
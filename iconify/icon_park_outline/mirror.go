package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mirror(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M34 44H14c-5 0-7-1-7-4v-6h11l1 4h11l1-4h10v6c0 3-2 4-7 4ZM12 34c-2.045-3.118-5-7-5-12.146C7 11.422 14.611 4 24 4s17 7.422 17 17.854C41 27 38.044 30.882 36 34"/>`),
		g.Group(children),
	)
}
package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortFromBottomToTopOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.763 3.289a.75.75 0 0 1 .837.261l3 4a.75.75 0 1 1-1.2.9l-1.65-2.2V20a.75.75 0 1 1-1.5 0V4a.75.75 0 0 1 .513-.711ZM13 8.75H4v-1.5h9v1.5Zm0 5H6v-1.5h7v1.5Zm0 5H8v-1.5h5v1.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
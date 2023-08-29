package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WickrMe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.212 13.582a2.068 2.068 0 0 1-1.816 1.068h0a2.143 2.143 0 0 1-2.137-2.137v-1.388a2.143 2.143 0 0 1 2.137-2.137h0a2.143 2.143 0 0 1 2.137 2.137v.747H19.26m-6.176-.704a2.178 2.178 0 0 1 2.17-2.171h0a2.178 2.178 0 0 1 2.172 2.17v3.474M8.74 8.997v5.645m0-3.474a2.178 2.178 0 0 1 2.171-2.171h0a2.178 2.178 0 0 1 2.171 2.17v3.474m10.067 22.62l14.784-14.784M13.436 37.261L28.22 22.477M40.5 5.501h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}
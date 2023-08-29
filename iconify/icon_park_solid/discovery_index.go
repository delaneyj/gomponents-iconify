package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscoveryIndex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDiscoveryIndex0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M16 6H8a2 2 0 0 0-2 2v8m10 26H8a2 2 0 0 1-2-2v-8m26 10h8a2 2 0 0 0 2-2v-8M32 6h8a2 2 0 0 1 2 2v8"/><path fill="#fff" d="M19 18h10v12H19z"/><path stroke-linecap="round" d="M24 18v-6m0 24v-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDiscoveryIndex0)"/>`),
		g.Group(children),
	)
}
package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTNetworkTree0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 34h8v8H4zM8 6h32v12H8z"/><path d="M24 34V18M8 34v-8h32v8"/><path fill="#555" d="M36 34h8v8h-8zm-16 0h8v8h-8z"/><path d="M14 12h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTNetworkTree0)"/>`),
		g.Group(children),
	)
}
package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudDirectLinkOneExchange(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 29H13c-1.1 0-2-.9-2-2v-4h2v4h14V13h-4v-2h4c1.1 0 2 .9 2 2v14c0 1.1-.9 2-2 2z"/><path fill="currentColor" d="M19 15v4h-4v2h6v-6zm-2-4h-6v6h2v-4h4z"/><path fill="currentColor" d="M5 3h14c1.1 0 2 .9 2 2v4h-2V5H5v14h4v2H5c-1.1 0-2-.9-2-2V5c0-1.1.9-2 2-2z"/>`),
		g.Group(children),
	)
}
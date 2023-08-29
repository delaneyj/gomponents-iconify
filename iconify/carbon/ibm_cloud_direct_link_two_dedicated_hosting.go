package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudDirectLinkTwoDedicatedHosting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 27H5V13h4v-2H5c-1.1 0-2 .9-2 2v6H0v2h3v6c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2v-4h-2v4z"/><path fill="currentColor" d="M11 19h10v2H11zm0-4h10v2H11zm0-4h10v2H11z"/><path fill="currentColor" d="M29 11V5c0-1.1-.9-2-2-2H13c-1.1 0-2 .9-2 2v4h2V5h14v14h-4v2h4c1.1 0 2-.9 2-2v-6h3v-2h-3z"/>`),
		g.Group(children),
	)
}
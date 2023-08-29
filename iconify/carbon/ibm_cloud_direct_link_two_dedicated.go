package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudDirectLinkTwoDedicated(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M32 11h-3V5c0-1.1-.9-2-2-2H13c-1.1 0-2 .9-2 2v4h2V5h14v14H13v-4h-2v4c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2v-6h3v-2z"/><path fill="currentColor" d="M21 17v-4c0-1.1-.9-2-2-2H5c-1.1 0-2 .9-2 2v6H0v2h3v6c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2v-4h-2v4H5V13h14v4h2z"/>`),
		g.Group(children),
	)
}
package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudDedicatedHost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="9" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="M26 2H6a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h9v4h2v-4h9a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm0 6H6V4h20zm-9 20v-2h-2v2H4v2h24v-2H17z"/><circle cx="9" cy="20" r="1" fill="currentColor"/><path fill="currentColor" d="M6 24h20a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2Zm0-6h20v4H6Z"/>`),
		g.Group(children),
	)
}
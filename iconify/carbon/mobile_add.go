package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 24h-4v-4h-2v4h-4v2h4v4h2v-4h4v-2z"/><path fill="currentColor" d="M10 28V10h12v7h2V6a2.002 2.002 0 0 0-2-2H10a2.002 2.002 0 0 0-2 2v22a2.002 2.002 0 0 0 2 2h6v-2Zm0-22h12v2H10Z"/>`),
		g.Group(children),
	)
}
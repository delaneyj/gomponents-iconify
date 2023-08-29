package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 7h4v4h-4zm0 6h4v4h-4zm-6-6h4v4h-4zm0 6h4v4h-4z"/><circle cx="14.5" cy="24.5" r="1.5" fill="currentColor"/><path fill="currentColor" d="M21 30H8a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h13v2H8v24h13v-8h2v8a2.002 2.002 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}
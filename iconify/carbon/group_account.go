package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupAccount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 11h-6a3 3 0 0 0-3 3v4h2v-4a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v1h2v-1a3 3 0 0 0-3-3zm-7-5a4 4 0 1 0 4-4a4 4 0 0 0-4 4zm6 0a2 2 0 1 1-2-2a2 2 0 0 1 2 2zm6 21h-6a2.002 2.002 0 0 1-2-2v-6a2.002 2.002 0 0 1 2-2h6a2.002 2.002 0 0 1 2 2v6a2.002 2.002 0 0 1-2 2zm-6-8v6h6v-6zM8 30H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h4v2H4v24h4zm20 0h-4v-2h4V4h-4V2h4a2.002 2.002 0 0 1 2 2v24a2.002 2.002 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}
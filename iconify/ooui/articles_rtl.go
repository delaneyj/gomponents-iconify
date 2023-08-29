package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArticlesRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 0a2 2 0 0 1 2 2H6a2 2 0 0 0-2 2v12a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M18 5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2zm-7 5H7V5h4zm5-4h-4V5h4zm0 2h-4V7h4zm0 2h-4V9h4zm0 2H7v-1h9zm0 2H7v-1h9zm0 2H7v-1h9z"/>`),
		g.Group(children),
	)
}
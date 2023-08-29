package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M18 16.634a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path fill-rule="evenodd" d="M5.866 3.134a1 1 0 1 0-1 1.732l13.454 7.768H2v4a4 4 0 0 0 4 4h12a4 4 0 0 0 4-4v-4l-16.134-9.5ZM20 14.634H4v2a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
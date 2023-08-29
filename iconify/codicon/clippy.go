package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clippy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 13.992H4v-9h8v2h1v-2.5l-.5-.5H11v-1h-1a2 2 0 0 0-4 0H4.94v1H3.5l-.5.5v10l.5.5H7v-1zm0-11.2a1 1 0 0 1 .8-.8a1 1 0 0 1 .58.06a.94.94 0 0 1 .45.36a1 1 0 1 1-1.75.94a1 1 0 0 1-.08-.56zm7.08 9.46L13 13.342v-5.35h-1v5.34l-1.08-1.08l-.71.71l1.94 1.93h.71l1.93-1.93l-.71-.71zm-5.92-4.16h.71l1.93 1.93l-.71.71l-1.08-1.08v5.34h-1v-5.35l-1.08 1.09l-.71-.71l1.94-1.93z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
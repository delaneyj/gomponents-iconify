package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M4.5 3a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 .5.5v14a.5.5 0 0 1-.855.352L10 12.676l-4.645 4.676A.5.5 0 0 1 4.5 17V3Zm1 .5v12.287l4.145-4.172a.5.5 0 0 1 .71 0l4.145 4.172V3.5h-9Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}
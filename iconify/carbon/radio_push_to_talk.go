package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioPushToTalk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 11h6v2h-6zm0 4h6v2h-6z"/><circle cx="16" cy="23" r="2" fill="currentColor"/><path fill="currentColor" d="M22 7h-1V2h-2v5h-9a2.002 2.002 0 0 0-2 2v2H6v2h2v2H6v2h2v11a2.002 2.002 0 0 0 2 2h12a2.002 2.002 0 0 0 2-2V9a2.002 2.002 0 0 0-2-2ZM10 28V9h12v19Z"/>`),
		g.Group(children),
	)
}
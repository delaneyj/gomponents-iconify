package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimpleClipboardEditor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.857 15.136h20.286m-6.091-6.584a4.052 4.052 0 0 0-8.104 0H9.298V43.5h29.403V8.552h-10.65Z"/><circle cx="24" cy="8.552" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}
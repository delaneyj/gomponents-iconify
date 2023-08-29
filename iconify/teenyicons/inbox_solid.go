package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5V8h-4.5a.5.5 0 0 0-.5.5a2.5 2.5 0 0 1-5 0a.5.5 0 0 0-.5-.5H0V1.5Z"/><path fill="currentColor" d="M0 9v4.5A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5V9h-4.035a3.5 3.5 0 0 1-6.93 0H0Z"/>`),
		g.Group(children),
	)
}
package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M14.5 20a2.5 2.5 0 0 1-5 0m13-2.5h-21v-.25l.848-.908a12 12 0 0 0 3.134-6.7l.336-2.684a6.23 6.23 0 0 1 12.364 0l.336 2.685a12 12 0 0 0 3.134 6.699l.848.908v.25Z"/>`),
		g.Group(children),
	)
}
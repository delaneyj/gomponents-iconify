package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Status(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.941 7.033a8 8 0 0 1-14.784 5.112a.75.75 0 1 1 1.283-.778a6.5 6.5 0 1 0 8.922-8.93a.75.75 0 0 1 .776-1.284a8 8 0 0 1 3.803 5.88ZM9 1a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM2.804 5a1 1 0 1 0-1.732-1a1 1 0 0 0 1.732 1ZM1 7a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm4-4.196a1 1 0 1 0-1-1.732a1 1 0 0 0 1 1.732Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
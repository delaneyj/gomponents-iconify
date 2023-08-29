package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.519 12a6.5 6.5 0 0 1 12.962 0h.269a4.75 4.75 0 1 1 0 9.5H7.25a4.75 4.75 0 1 1 0-9.5h.269ZM14 7.5a5 5 0 0 0-5 5v.25a.75.75 0 0 1-.75.75h-1a3.25 3.25 0 0 0 0 6.5h13.5a3.25 3.25 0 0 0 0-6.5h-1a.75.75 0 0 1-.75-.75v-.25a5 5 0 0 0-5-5Z"/>`),
		g.Group(children),
	)
}
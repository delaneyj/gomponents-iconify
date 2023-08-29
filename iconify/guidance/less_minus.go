package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LessMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 14.5h23v-.25l-.054-.27a10.101 10.101 0 0 1 0-3.96l.054-.27V9.5H.5v.25l.054.27a10.1 10.1 0 0 1 0 3.96l-.054.27v.25Z"/>`),
		g.Group(children),
	)
}
package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Confused(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 9a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm5.66 4.56l-4.19 1.5A1 1 0 0 0 10.8 17a1 1 0 0 0 .34-.06l4.2-1.5a1 1 0 1 0-.68-1.88ZM15 9a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-3-7a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}
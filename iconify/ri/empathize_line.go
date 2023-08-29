package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmpathizeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.364 10.98a4 4 0 0 1 0 5.656l-5.657 5.657a1 1 0 0 1-1.414 0l-5.657-5.657a4 4 0 1 1 5.657-5.657l.707.707l.707-.707a4 4 0 0 1 5.657 0ZM7.051 12.392a2 2 0 0 0 0 2.829l4.95 4.95l4.95-4.95a2 2 0 1 0-2.83-2.827l-2.123 2.118l-2.119-2.12a2 2 0 0 0-2.828 0ZM12 1a4 4 0 1 1 0 8a4 4 0 0 1 0-8Zm0 2a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		g.Group(children),
	)
}
package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonHearts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 1.246c.832-.855 2.913.642 0 2.566c-2.913-1.924-.832-3.421 0-2.566ZM9 5a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-9 8c0 1 1 1 1 1h10s1 0 1-1s-1-4-6-4s-6 3-6 4Zm13.5-8.09c1.387-1.425 4.855 1.07 0 4.277c-4.854-3.207-1.387-5.702 0-4.276ZM15 2.165c.555-.57 1.942.428 0 1.711c-1.942-1.283-.555-2.281 0-1.71Z"/>`),
		g.Group(children),
	)
}
package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SynchronizeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20.5 3a1 1 0 0 0-1.707-.707l-1.33 1.33A9.955 9.955 0 0 0 12 2C6.477 2 2 6.477 2 12a1 1 0 0 0 2 0a8 8 0 0 1 12.01-6.924l-1.217 1.217A1 1 0 0 0 15.5 8h4a1 1 0 0 0 1-1V3ZM7.99 18.924l1.217-1.217A1 1 0 0 0 8.5 16h-4a1 1 0 0 0-1 1v4a1 1 0 0 0 1.707.707l1.33-1.33A9.956 9.956 0 0 0 12 22c5.523 0 10-4.477 10-10a1 1 0 1 0-2 0a8 8 0 0 1-12.01 6.924Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundCarpenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.73 14.23L7.71 2.21a.996.996 0 0 0-1.41 0L3.7 4.8a1 1 0 0 0-.11 1.28l7.65 10.98c-.78.78-.78 2.05 0 2.83l1.41 1.41c.78.78 2.05.78 2.83 0l4.24-4.24c.79-.78.79-2.05.01-2.83zm-5.66 5.65l-1.41-1.41l4.24-4.24l1.41 1.41l-4.24 4.24z"/>`),
		g.Group(children),
	)
}
package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundRemoveDone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.14 2.69a.996.996 0 0 0 0 1.41l9.67 9.67l-1.41 1.41l-3.54-3.53a.996.996 0 1 0-1.41 1.41l4.24 4.24c.39.39 1.02.39 1.41 0l2.12-2.12l5.89 5.89a.996.996 0 1 0 1.41-1.41L5.55 2.69a.996.996 0 0 0-1.41 0zm13.91 9.67l4.24-4.24a.999.999 0 0 0-.01-1.42c-.39-.38-1.02-.38-1.41.01l-4.24 4.24l1.42 1.41zM16.64 6.7a.996.996 0 0 0-1.41 0l-1.42 1.42l1.41 1.41l1.42-1.42a.996.996 0 0 0 0-1.41zM1.79 13.06l4.95 4.95l1.41-1.41l-4.95-4.95a.996.996 0 1 0-1.41 1.41z"/>`),
		g.Group(children),
	)
}
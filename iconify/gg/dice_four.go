package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16.945 5.055a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-2 11.837a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm-7.89-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-2-7.837a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z"/><path fill-rule="evenodd" d="M4 1a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3H4Zm16 2H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
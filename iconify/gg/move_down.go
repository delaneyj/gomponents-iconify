package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7 5h2v8H7V5Zm8 0h2v8h-2V5Z"/><path d="M11 5h2v10h3.035l-4 4.071l-4-4.071H11V5Z"/></g>`),
		g.Group(children),
	)
}
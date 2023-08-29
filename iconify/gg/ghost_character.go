package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GhostCharacter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10.976 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm2.995 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path fill-rule="evenodd" d="M19 21V10a7 7 0 1 0-14 0v11h2.828l1.415-1.414L10.657 21h2.686l1.414-1.414L16.172 21H19Zm-2-11a5 5 0 0 0-10 0v9l2.243-2.243L12 19.515l2.757-2.758L17 19v-9Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
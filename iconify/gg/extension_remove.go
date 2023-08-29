package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExtensionRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 11V5H4v14h14v-8h-6ZM6 7h4v4H6V7Zm4 6v4H6v-4h4Zm6 0v4h-4v-4h4Z" clip-rule="evenodd"/><path d="M20 7h-6v2h6V7Z"/></g>`),
		g.Group(children),
	)
}
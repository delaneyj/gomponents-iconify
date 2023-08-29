package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExtensionAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16 4h2v2h2v2h-2v2h-2V8h-2V6h2V4Z"/><path fill-rule="evenodd" d="M12 12V6H4v14h14v-8h-6ZM6 8h4v4H6V8Zm4 6v4H6v-4h4Zm6 0v4h-4v-4h4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
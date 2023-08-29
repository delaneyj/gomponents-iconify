package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5 5a3 3 0 0 1 5-2.236A3 3 0 0 1 14.83 6H16a2 2 0 1 1 0 4h-5V9a1 1 0 1 0-2 0v1H4a2 2 0 1 1 0-4h1.17C5.06 5.687 5 5.35 5 5Zm4 1V5a1 1 0 1 0-1 1h1Zm3 0a1 1 0 1 0-1-1v1h1Z" clip-rule="evenodd"/><path d="M9 11H3v5a2 2 0 0 0 2 2h4v-7Zm2 7h4a2 2 0 0 0 2-2v-5h-6v7Z"/></g>`),
		g.Group(children),
	)
}
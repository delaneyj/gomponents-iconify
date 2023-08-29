package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M5 3a1 1 0 0 0 0 2c5.523 0 10 4.477 10 10a1 1 0 1 0 2 0C17 8.373 11.627 3 5 3Z"/><path d="M4 9a1 1 0 0 1 1-1a7 7 0 0 1 7 7a1 1 0 1 1-2 0a5 5 0 0 0-5-5a1 1 0 0 1-1-1Zm-1 6a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z"/></g>`),
		g.Group(children),
	)
}
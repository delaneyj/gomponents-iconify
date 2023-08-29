package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M9 7a1 1 0 0 1 2 0v4a1 1 0 1 1-2 0V7Z"/><circle cx="10" cy="14" r="1"/><path fill-rule="evenodd" d="M2 10a8 8 0 1 0 16 0a8 8 0 0 0-16 0Zm14 0a6 6 0 1 1-12 0a6 6 0 0 1 12 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
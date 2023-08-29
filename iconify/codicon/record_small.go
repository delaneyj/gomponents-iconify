package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M8 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M12 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm-1 0a3 3 0 1 0-6 0a3 3 0 0 0 6 0Z"/></g>`),
		g.Group(children),
	)
}
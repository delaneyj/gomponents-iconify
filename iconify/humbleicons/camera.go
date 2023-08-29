package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path d="M20 17V8a1 1 0 0 0-1-1h-2.382a1 1 0 0 1-.894-.553l-.448-.894A1 1 0 0 0 14.382 5H9.618a1 1 0 0 0-.894.553l-.448.894A1 1 0 0 1 7.382 7H5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1Z"/><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/></g>`),
		g.Group(children),
	)
}
package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.438 6.062h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1Zm0 6.438h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1Zm0 6.435h-9a.5.5 0 1 1 0-1h9a.5.5 0 0 1 0 1ZM5.562 8.062a2.5 2.5 0 1 1 2.5-2.5a2.5 2.5 0 0 1-2.5 2.5Zm0-4a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5Zm0 10.438a2.5 2.5 0 1 1 2.5-2.5a2.5 2.5 0 0 1-2.5 2.5Zm0-4a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5Zm0 10.438a2.5 2.5 0 1 1 2.5-2.5a2.5 2.5 0 0 1-2.5 2.5Zm0-4a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		g.Group(children),
	)
}
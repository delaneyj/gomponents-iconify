package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TentTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.601 3.211a.75.75 0 0 0-1.152 0c-1.795 2.154-5.337 4.71-7.374 5.848a.75.75 0 0 0-.377.552L4.06 21.5H2.75a.75.75 0 0 0 0 1.5h22.5a.75.75 0 0 0 0-1.5h-1.26L22.352 9.61a.75.75 0 0 0-.378-.552c-2.036-1.138-5.578-3.694-7.373-5.848ZM10.272 21.5c1.259-1.83 2.557-4.18 3.728-7.892c1.059 3.35 2.24 5.665 3.728 7.892h-7.456Z"/>`),
		g.Group(children),
	)
}
package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.44 15.94a2.5 2.5 0 0 0-1.96.95l-8.51-4.25a2.356 2.356 0 0 0 0-1.29l8.5-4.25a2.5 2.5 0 1 0-.53-1.54a2.269 2.269 0 0 0 .09.65l-8.5 4.25a2.5 2.5 0 1 0 0 3.08l8.5 4.25a2.269 2.269 0 0 0-.09.65a2.5 2.5 0 1 0 2.5-2.5Zm0-11.88a1.5 1.5 0 1 1-1.5 1.5a1.5 1.5 0 0 1 1.5-1.5ZM5.56 13.5a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5Zm12.88 6.44a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5Z"/>`),
		g.Group(children),
	)
}
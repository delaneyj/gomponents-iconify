package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lightning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M240 0L5 299l128 21l-42 192l256-320l-128-21zm28 222l-96 119l2-12l11-45l-45-7l-58-8l96-122l-2 19l-4 41l40 6z"/>`),
		g.Group(children),
	)
}
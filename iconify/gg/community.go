package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Community(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 17.456a6 6 0 1 1 0-10.912a6 6 0 1 1 0 10.912Zm-2-1.487a4 4 0 1 1 0-7.938A5.977 5.977 0 0 0 8.5 12a5.98 5.98 0 0 0 1.5 3.969Zm4-7.938a4 4 0 1 1 0 7.938A5.978 5.978 0 0 0 15.5 12A5.978 5.978 0 0 0 14 8.031Zm-2 .846c.915.733 1.5 1.86 1.5 3.123c0 1.263-.585 2.39-1.5 3.123A3.993 3.993 0 0 1 10.5 12c0-1.263.585-2.39 1.5-3.123Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
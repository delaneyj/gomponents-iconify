package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PassportOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 21.936h-9a2.5 2.5 0 0 1-2.5-2.5V4.564a2.5 2.5 0 0 1 2.5-2.5h9a2.5 2.5 0 0 1 2.5 2.5v14.872a2.5 2.5 0 0 1-2.5 2.5Zm-9-18.872a1.5 1.5 0 0 0-1.5 1.5v14.872a1.5 1.5 0 0 0 1.5 1.5h9a1.5 1.5 0 0 0 1.5-1.5V4.564a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M12 12.563a3.5 3.5 0 1 1 3.5-3.5a3.5 3.5 0 0 1-3.5 3.5Zm0-6a2.5 2.5 0 1 0 2.5 2.5a2.5 2.5 0 0 0-2.5-2.5Zm3 11.875H9a.5.5 0 0 1 0-1h6a.5.5 0 1 1 0 1Z"/>`),
		g.Group(children),
	)
}
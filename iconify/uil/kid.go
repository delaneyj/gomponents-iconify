package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 10a1 1 0 1 0-1 1a1 1 0 0 0 1-1Zm4.5 4.06a1 1 0 0 0-1.37.36a1.3 1.3 0 0 1-2.26 0a1 1 0 0 0-1.37-.36a1 1 0 0 0-.37 1.36a3.31 3.31 0 0 0 5.74 0a1 1 0 0 0-.37-1.36ZM15 9a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-3-7a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18A8 8 0 0 1 9 4.57A3 3 0 0 0 9 5a3 3 0 0 0 3 3a1 1 0 0 0 0-2a1 1 0 0 1 0-2a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}
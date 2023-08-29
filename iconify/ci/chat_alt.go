package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 2a2 2 0 0 1 2 2H4v11.177a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h12Z"/><path fill="currentColor" d="m14 22l-2.667-2.823H8a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v9.177a2 2 0 0 1-2 2h-3.333L14 22Zm1.805-4.823H20V8H8v9.177h4.195L14 19.087l1.805-1.91Z"/>`),
		g.Group(children),
	)
}
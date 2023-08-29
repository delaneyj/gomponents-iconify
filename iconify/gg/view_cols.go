package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCols(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V8Zm14-1h3a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-3V7Zm-2 0h-4v10h4V7ZM8 17V7H5a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSixFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.532 3.153a1 1 0 0 1 .315 1.379l-2.858 4.553a6 6 0 1 1-4.05 2.691l5.214-8.308a1 1 0 0 1 1.379-.315ZM12 11a4 4 0 1 0 0 8a4 4 0 0 0 0-8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
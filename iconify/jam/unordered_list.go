package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnorderedList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 0h9a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2zm0 8h9a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2zm0-4h9a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2zM1 0a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm0 8a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm0-4a1 1 0 1 1 0 2a1 1 0 0 1 0-2z"/>`),
		g.Group(children),
	)
}
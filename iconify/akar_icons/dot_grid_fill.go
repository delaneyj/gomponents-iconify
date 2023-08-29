package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotGridFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm8 0a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm10 2a2 2 0 1 0-4 0a2 2 0 0 0 4 0ZM4 10a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm10 2a2 2 0 1 0-4 0a2 2 0 0 0 4 0Zm6-2a2 2 0 1 1 0 4a2 2 0 0 1 0-4ZM6 20a2 2 0 1 0-4 0a2 2 0 0 0 4 0Zm6-2a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm10 2a2 2 0 1 0-4 0a2 2 0 0 0 4 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
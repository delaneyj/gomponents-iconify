package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridRound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 20a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm12-6a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm12-6a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4ZM6 8a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}
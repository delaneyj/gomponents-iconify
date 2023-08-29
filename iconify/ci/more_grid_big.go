package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreGridBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 18a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm12-6a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm12-6a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0ZM5 6a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/>`),
		g.Group(children),
	)
}
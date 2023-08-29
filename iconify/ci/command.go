package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Command(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 15v3a3 3 0 1 1-3-3h3Zm0 0h6m-6 0V9m6 6v3a3 3 0 1 0 3-3h-3Zm0 0V9m0 0H9m6 0V6a3 3 0 1 1 3 3h-3ZM9 9V6a3 3 0 1 0-3 3h3Z"/>`),
		g.Group(children),
	)
}
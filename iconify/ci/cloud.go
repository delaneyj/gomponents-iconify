package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11a4 4 0 0 1 0 8H6a5 5 0 0 1-.331-9.99A7.002 7.002 0 0 1 18.929 11H19Z"/>`),
		g.Group(children),
	)
}
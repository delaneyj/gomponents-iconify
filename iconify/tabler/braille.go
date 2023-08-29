package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Braille(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 5a1 1 0 1 0 2 0a1 1 0 0 0-2 0zM7 5a1 1 0 1 0 2 0a1 1 0 0 0-2 0zm0 14a1 1 0 1 0 2 0a1 1 0 0 0-2 0zm9-7h.01M8 12h.01M16 19h.01"/>`),
		g.Group(children),
	)
}
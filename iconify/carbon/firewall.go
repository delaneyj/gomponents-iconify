package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Firewall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 20.397v3c-1 0-2-1.5-2-4v-3c-4 5-5 7-5 9A5 5 0 0 0 23.046 30A7.528 7.528 0 0 1 25 26.397A7.528 7.528 0 0 1 26.954 30A5 5 0 0 0 30 25.397c0-2-1.125-3.571-3-5zM17 28H4v-4h13v-2H4a2.002 2.002 0 0 0-2 2v4a2.002 2.002 0 0 0 2 2h13z"/><path fill="currentColor" d="M28 12H7a2.002 2.002 0 0 0-2 2v4a2.002 2.002 0 0 0 2 2h10v-2H7v-4h21l.001 2H30v-2a2.002 2.002 0 0 0-2-2zm-3-2H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h21a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zM4 4v4h21V4z"/>`),
		g.Group(children),
	)
}
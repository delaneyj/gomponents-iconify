package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16v-3m0 0v-3m0 3H9m3 0h3m8 2a4 4 0 0 0-4.07-4A7.002 7.002 0 0 0 5.669 9.01A5 5 0 0 0 6 19h13a4 4 0 0 0 4-4Z"/>`),
		g.Group(children),
	)
}
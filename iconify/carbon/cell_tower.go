package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellTower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 11v5h-8v-5h-2v5H7v-5H5v12h2v-5h3v12h2V18h3v5h2v-5h3v12h2V18h3v5h2V11zm-9-5c-1.7 0-3.2.7-4.2 1.8l1.4 1.4C13.9 8.4 14.9 8 16 8s2.1.4 2.8 1.2l1.4-1.4C19.2 6.7 17.7 6 16 6z"/><path fill="currentColor" d="m8.9 4.9l1.4 1.4C11.8 4.9 13.8 4 16 4s4.2.9 5.7 2.3l1.4-1.4C21.3 3.1 18.8 2 16 2s-5.3 1.1-7.1 2.9z"/>`),
		g.Group(children),
	)
}
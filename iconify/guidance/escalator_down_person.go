package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorDownPerson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5 5.5l9 9h14v9H10.642a10 10 0 0 1-7.07-2.929L.5 17.5m11-3l1.018-5.588A3 3 0 0 1 15.459 6.5H16l.598 2.99a2.5 2.5 0 0 0 2.451 2.01H22m-5.65-7s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}
package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneOclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm4 30c0 1.477-.81 2.752-2 3.445V38h-4v-.314l-3-1.613l1.427-2.308A3.97 3.97 0 0 1 28 32c0-1.477.81-2.752 2-3.445V6h4v18.756L39.416 16L43 17.934l-7.513 12.143c.317.573.513 1.222.513 1.923z"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}
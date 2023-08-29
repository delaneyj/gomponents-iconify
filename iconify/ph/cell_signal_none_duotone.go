package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellSignalNoneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M198.12 25.23a16 16 0 0 0-17.43 3.47l-160 160A16 16 0 0 0 32 216h160a16 16 0 0 0 16-16V40a16 16 0 0 0-9.88-14.77ZM192 200H32L192 40Z"/>`),
		g.Group(children),
	)
}
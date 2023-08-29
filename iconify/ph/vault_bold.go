package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VaultBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 36H40a20 20 0 0 0-20 20v136a20 20 0 0 0 20 20h12v12a12 12 0 0 0 24 0v-12h104v12a12 12 0 0 0 24 0v-12h12a20 20 0 0 0 20-20V56a20 20 0 0 0-20-20ZM44 188V60h168v56h-21.68a44 44 0 1 0 0 24H212v48Zm124-60a20 20 0 1 1-20-20a20 20 0 0 1 20 20Z"/>`),
		g.Group(children),
	)
}
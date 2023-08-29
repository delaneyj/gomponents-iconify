package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircuitryBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20ZM52 52h24v100a20 20 0 1 0 24 0v-27l40 40v39H52Zm152 152h-40v-44a12 12 0 0 0-3.51-8.49L100 91V52h24v20a12 12 0 0 0 3.51 8.49l20.71 20.7a20.17 20.17 0 1 0 17-17L148 67V52h56Z"/>`),
		g.Group(children),
	)
}
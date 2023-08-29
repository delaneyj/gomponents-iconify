package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorScreenBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 72V48a20 20 0 0 0-20-20H40a20 20 0 0 0-20 20v24a20 20 0 0 0 16 19.6V164h-4a12 12 0 0 0 0 24h84v23.22a24 24 0 1 0 24 0V188h84a12 12 0 0 0 0-24h-4V91.6A20 20 0 0 0 236 72ZM44 52h168v16H44Zm16 112V92h136v72Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 72H56a8 8 0 0 1 0-16h136a8 8 0 0 0 0-16H56a24 24 0 0 0-24 24v128a24 24 0 0 0 24 24h160a16 16 0 0 0 16-16V88a16 16 0 0 0-16-16Zm0 128H56a8 8 0 0 1-8-8V86.63A23.84 23.84 0 0 0 56 88h160Zm-48-60a12 12 0 1 1 12 12a12 12 0 0 1-12-12Z"/>`),
		g.Group(children),
	)
}
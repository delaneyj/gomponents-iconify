package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlackLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M214.4 128A28 28 0 1 0 180 84.43V56a28 28 0 0 0-52-14.4A28 28 0 1 0 84.43 76H56a28 28 0 0 0-14.4 52A28 28 0 1 0 76 171.57V200a28 28 0 0 0 52 14.41A28 28 0 1 0 171.57 180H200a28 28 0 0 0 14.4-52ZM76 152a20 20 0 1 1-20-20h20Zm48 48a20 20 0 0 1-40 0v-48a20 20 0 0 1 20-20h20Zm0-76H56a20 20 0 0 1 0-40h48a20 20 0 0 1 20 20Zm0-48h-20a20 20 0 1 1 20-20Zm56 28a20 20 0 1 1 20 20h-20Zm-48-48a20 20 0 0 1 40 0v48a20 20 0 0 1-20 20h-20Zm40 144a20 20 0 0 1-40 0v-20h20a20 20 0 0 1 20 20Zm28-28h-48a20 20 0 0 1-20-20v-20h68a20 20 0 0 1 0 40Z"/>`),
		g.Group(children),
	)
}
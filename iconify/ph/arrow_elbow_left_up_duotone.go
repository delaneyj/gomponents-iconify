package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowLeftUpDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M128 96H32l48-48Z" opacity=".2"/><path d="M224 184H88v-80h40a8 8 0 0 0 5.66-13.66l-48-48a8 8 0 0 0-11.32 0l-48 48A8 8 0 0 0 32 104h40v88a8 8 0 0 0 8 8h144a8 8 0 0 0 0-16ZM80 59.31L108.69 88H51.31Z"/></g>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowDownLeftDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M96 128v96l-48-48Z" opacity=".2"/><path d="M192 24a8 8 0 0 0-8 8v136h-80v-40a8 8 0 0 0-13.66-5.66l-48 48a8 8 0 0 0 0 11.32l48 48A8 8 0 0 0 104 224v-40h88a8 8 0 0 0 8-8V32a8 8 0 0 0-8-8ZM88 204.69L59.31 176L88 147.31Z"/></g>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChalkboardSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M184 168v32h-64v-32Z" opacity=".2"/><path d="M24 168V56a16 16 0 0 1 16-16h176a16 16 0 0 1 16 16v112a8 8 0 0 1-16 0V56H40v112a8 8 0 0 1-16 0Zm224 32a8 8 0 0 1-8 8H16a8 8 0 0 1 0-16h96v-24a8 8 0 0 1 8-8h64a8 8 0 0 1 8 8v24h48a8 8 0 0 1 8 8Zm-120-8h48v-16h-48Z"/></g>`),
		g.Group(children),
	)
}
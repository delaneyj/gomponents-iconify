package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChalkboardSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M20 160V56a20 20 0 0 1 20-20h176a20 20 0 0 1 20 20v104a12 12 0 0 1-24 0V60H44v100a12 12 0 0 1-24 0Zm232 40a12 12 0 0 1-12 12H16a12 12 0 0 1 0-24h92v-28a12 12 0 0 1 12-12h64a12 12 0 0 1 12 12v28h44a12 12 0 0 1 12 12Zm-120-12h40v-16h-40Z"/>`),
		g.Group(children),
	)
}
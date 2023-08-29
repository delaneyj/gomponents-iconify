package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarageBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 188h-4V98.67A20 20 0 0 0 227.09 82l-88-58.66a19.94 19.94 0 0 0-22.18 0L28.91 82A20 20 0 0 0 20 98.67V188h-4a12 12 0 0 0 0 24h224a12 12 0 0 0 0-24ZM44 100.81l84-56l84 56V188h-16v-68a12 12 0 0 0-12-12H72a12 12 0 0 0-12 12v68H44ZM172 132v16h-32v-16Zm-56 16H84v-16h32Zm-32 24h32v16H84Zm56 0h32v16h-32Z"/>`),
		g.Group(children),
	)
}
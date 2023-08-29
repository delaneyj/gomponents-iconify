package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuitcaseBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 48h-36v-8a28 28 0 0 0-28-28h-48a28 28 0 0 0-28 28v8H40a20 20 0 0 0-20 20v128a20 20 0 0 0 20 20h176a20 20 0 0 0 20-20V68a20 20 0 0 0-20-20Zm-116-8a4 4 0 0 1 4-4h48a4 4 0 0 1 4 4v8h-56Zm56 32v120h-56V72ZM44 72h32v120H44Zm168 120h-32V72h32Z"/>`),
		g.Group(children),
	)
}
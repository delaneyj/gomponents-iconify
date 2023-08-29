package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M232 64.49v127a8.49 8.49 0 0 1-8.49 8.49h-191a8.49 8.49 0 0 1-8.51-8.47v-127A8.49 8.49 0 0 1 32.49 56h191a8.49 8.49 0 0 1 8.51 8.49Z" opacity=".2"/><path d="M223.51 48h-191A16.51 16.51 0 0 0 16 64.49v127A16.51 16.51 0 0 0 32.49 208h191A16.51 16.51 0 0 0 240 191.51v-127A16.51 16.51 0 0 0 223.51 48Zm.49 143.51a.49.49 0 0 1-.49.49h-191a.49.49 0 0 1-.49-.49v-127a.49.49 0 0 1 .49-.49h191a.49.49 0 0 1 .49.49ZM208 128a8 8 0 0 1-8 8H56a8 8 0 0 1 0-16h144a8 8 0 0 1 8 8Zm0-32a8 8 0 0 1-8 8H56a8 8 0 0 1 0-16h144a8 8 0 0 1 8 8ZM72 160a8 8 0 0 1-8 8h-8a8 8 0 0 1 0-16h8a8 8 0 0 1 8 8Zm96 0a8 8 0 0 1-8 8H96a8 8 0 0 1 0-16h64a8 8 0 0 1 8 8Zm40 0a8 8 0 0 1-8 8h-8a8 8 0 0 1 0-16h8a8 8 0 0 1 8 8Z"/></g>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M223.51 52h-191A12.5 12.5 0 0 0 20 64.49v127A12.5 12.5 0 0 0 32.49 204h191A12.5 12.5 0 0 0 236 191.51v-127A12.5 12.5 0 0 0 223.51 52ZM228 191.51a4.49 4.49 0 0 1-4.49 4.49h-191a4.49 4.49 0 0 1-4.51-4.49v-127A4.49 4.49 0 0 1 32.49 60h191a4.49 4.49 0 0 1 4.51 4.49ZM204 128a4 4 0 0 1-4 4H56a4 4 0 0 1 0-8h144a4 4 0 0 1 4 4Zm0-32a4 4 0 0 1-4 4H56a4 4 0 0 1 0-8h144a4 4 0 0 1 4 4ZM68 160a4 4 0 0 1-4 4h-8a4 4 0 0 1 0-8h8a4 4 0 0 1 4 4Zm96 0a4 4 0 0 1-4 4H96a4 4 0 0 1 0-8h64a4 4 0 0 1 4 4Zm40 0a4 4 0 0 1-4 4h-8a4 4 0 0 1 0-8h8a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}
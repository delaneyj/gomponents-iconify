package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M223.51 50h-191A14.51 14.51 0 0 0 18 64.49v127A14.51 14.51 0 0 0 32.49 206h191A14.51 14.51 0 0 0 238 191.51v-127A14.51 14.51 0 0 0 223.51 50ZM226 191.51a2.49 2.49 0 0 1-2.49 2.49h-191a2.49 2.49 0 0 1-2.51-2.49v-127A2.49 2.49 0 0 1 32.49 62h191a2.49 2.49 0 0 1 2.51 2.49ZM206 128a6 6 0 0 1-6 6H56a6 6 0 0 1 0-12h144a6 6 0 0 1 6 6Zm0-32a6 6 0 0 1-6 6H56a6 6 0 0 1 0-12h144a6 6 0 0 1 6 6ZM70 160a6 6 0 0 1-6 6h-8a6 6 0 0 1 0-12h8a6 6 0 0 1 6 6Zm96 0a6 6 0 0 1-6 6H96a6 6 0 0 1 0-12h64a6 6 0 0 1 6 6Zm40 0a6 6 0 0 1-6 6h-8a6 6 0 0 1 0-12h8a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}
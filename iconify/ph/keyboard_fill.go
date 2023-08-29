package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M223.51 48h-191A16.51 16.51 0 0 0 16 64.49v127A16.51 16.51 0 0 0 32.49 208h191A16.51 16.51 0 0 0 240 191.51v-127A16.51 16.51 0 0 0 223.51 48ZM64 168H48a8 8 0 0 1 0-16h16a8 8 0 0 1 0 16Zm96 0H96a8 8 0 0 1 0-16h64a8 8 0 0 1 0 16Zm48 0h-16a8 8 0 0 1 0-16h16a8 8 0 0 1 0 16Zm0-32H48a8 8 0 0 1 0-16h160a8 8 0 0 1 0 16Zm0-32H48a8 8 0 0 1 0-16h160a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}
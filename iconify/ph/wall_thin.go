package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 52H32a4 4 0 0 0-4 4v144a4 4 0 0 0 4 4h192a4 4 0 0 0 4-4V56a4 4 0 0 0-4-4ZM84 148v-40h88v40Zm-48 0v-40h40v40Zm144-40h40v40h-40Zm40-8h-88V60h88Zm-96-40v40H36V60Zm-88 96h88v40H36Zm96 40v-40h88v40Z"/>`),
		g.Group(children),
	)
}
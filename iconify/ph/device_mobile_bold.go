package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceMobileBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 12H80a28 28 0 0 0-28 28v176a28 28 0 0 0 28 28h96a28 28 0 0 0 28-28V40a28 28 0 0 0-28-28ZM76 76h104v104H76Zm4-40h96a4 4 0 0 1 4 4v12H76V40a4 4 0 0 1 4-4Zm96 184H80a4 4 0 0 1-4-4v-12h104v12a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}
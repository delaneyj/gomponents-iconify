package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceMobileThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 20H80a20 20 0 0 0-20 20v176a20 20 0 0 0 20 20h96a20 20 0 0 0 20-20V40a20 20 0 0 0-20-20ZM68 60h120v136H68Zm12-32h96a12 12 0 0 1 12 12v12H68V40a12 12 0 0 1 12-12Zm96 200H80a12 12 0 0 1-12-12v-12h120v12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}
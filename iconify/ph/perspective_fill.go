package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PerspectiveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 120H32V77.09a16 16 0 0 1 13.14-15.74l160-29.09A16 16 0 0 1 224 48ZM32 178.91a16 16 0 0 0 13.14 15.74l160 29.09a16.47 16.47 0 0 0 2.86.26a16 16 0 0 0 16-16v-72H32ZM240 120h-16v16h16a8 8 0 0 0 0-16Zm-224 0a8 8 0 0 0 0 16h16v-16Z"/>`),
		g.Group(children),
	)
}
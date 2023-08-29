package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenSmartphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M104 48v416a32.036 32.036 0 0 0 32 32h248a32.036 32.036 0 0 0 32-32V48a32.036 32.036 0 0 0-32-32H136a32.036 32.036 0 0 0-32 32Zm280.021 416H136V48h248Z"/><path fill="currentColor" d="M216 80h96v32h-96zm32 312h32v32h-32z"/>`),
		g.Group(children),
	)
}
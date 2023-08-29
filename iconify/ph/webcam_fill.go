package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebcamFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 104a32 32 0 1 1-32-32a32 32 0 0 1 32 32Zm72 104a8 8 0 0 1-8 8H32a8 8 0 0 1 0-16h88v-16.4a80 80 0 1 1 16 0V200h88a8 8 0 0 1 8 8Zm-104-56a48 48 0 1 0-48-48a48.05 48.05 0 0 0 48 48Z"/>`),
		g.Group(children),
	)
}
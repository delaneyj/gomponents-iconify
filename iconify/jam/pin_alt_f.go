package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinAltF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 10.465a5.002 5.002 0 0 1 1-9.9a5 5 0 0 1 1 9.9v9.1a1 1 0 0 1-2 0v-9.1z"/>`),
		g.Group(children),
	)
}
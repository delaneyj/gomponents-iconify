package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.793 7.5L1.146 1.854l.708-.708L7.5 6.793l5.646-5.647l.708.708L8.207 7.5l5.647 5.646l-.707.707L7.5 8.207l-5.646 5.646l-.708-.707L6.793 7.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
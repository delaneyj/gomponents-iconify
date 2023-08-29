package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 0A4.5 4.5 0 0 0 7 8.973V15h1V8.973A4.5 4.5 0 0 0 7.5 0Z"/>`),
		g.Group(children),
	)
}
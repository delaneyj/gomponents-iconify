package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.707.003a.5.5 0 0 0-.375.846a6 6 0 0 1-5.569 10.024a.5.5 0 0 0-.519.765A7.5 7.5 0 1 0 7.707.003Z"/>`),
		g.Group(children),
	)
}
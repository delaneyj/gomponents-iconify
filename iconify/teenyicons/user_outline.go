package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M10.5 3.498a2.999 2.999 0 0 1-3 2.998a2.999 2.999 0 1 1 3-2.998Zm2 10.992h-10v-1.996a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v1.997Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
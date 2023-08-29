package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M4.5 7.5h4M11 4L8.5 7.495L11 11m3.5-8.501a2 2 0 0 1-4 0a2 2 0 0 1 4 0Zm0 9.993a2 2 0 0 1-4 0a2 2 0 0 1 4 0Zm-10-4.997a2 2 0 0 1-4 0a2 2 0 0 1 4 0Z"/>`),
		g.Group(children),
	)
}
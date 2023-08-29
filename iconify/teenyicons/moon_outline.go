package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M1.66 11.362A6.5 6.5 0 0 0 7.693.502a7 7 0 1 1-6.031 10.86Z"/>`),
		g.Group(children),
	)
}
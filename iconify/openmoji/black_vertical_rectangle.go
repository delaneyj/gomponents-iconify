package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackVerticalRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M55 67V5H17v62h38Z"/><path fill="#3F3F3F" d="M55 67V5H17v62h38Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M55 67V5H17v62h38Z"/>`),
		g.Group(children),
	)
}
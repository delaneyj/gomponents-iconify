package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomInSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 6.5a6.5 6.5 0 1 1 11.436 4.23l3.418 3.416l-.707.708l-3.418-3.418A6.5 6.5 0 0 1 0 6.5ZM6 9V7H4V6h2V4h1v2h2v1H7v2H6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
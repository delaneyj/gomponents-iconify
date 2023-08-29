package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsThreeVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M140 128a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm-12-56a12 12 0 1 0-12-12a12 12 0 0 0 12 12Zm0 112a12 12 0 1 0 12 12a12 12 0 0 0-12-12Z"/>`),
		g.Group(children),
	)
}
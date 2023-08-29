package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 64h480v32H16zm0 88h328v32H16zm0 88h480v32H16zm0 88h328v32H16zm0 88h480v32H16z"/>`),
		g.Group(children),
	)
}
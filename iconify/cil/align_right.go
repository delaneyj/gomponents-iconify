package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 64h480v32H16zm152 88h328v32H168zM16 240h480v32H16zm152 88h328v32H168zM16 416h480v32H16z"/>`),
		g.Group(children),
	)
}
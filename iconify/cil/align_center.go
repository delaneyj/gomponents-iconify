package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 64h480v32H16zm80 88h320v32H96zm-80 88h480v32H16zm80 88h320v32H96zm-80 88h480v32H16z"/>`),
		g.Group(children),
	)
}
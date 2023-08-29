package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 56v144a12 12 0 0 1-24 0V56a12 12 0 0 1 24 0Zm-40 40v64a20 20 0 0 1-20 20H32a20 20 0 0 1-20-20V96a20 20 0 0 1 20-20h144a20 20 0 0 1 20 20Zm-24 4H36v56h136Z"/>`),
		g.Group(children),
	)
}
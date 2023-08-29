package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleEightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm0 192a84 84 0 1 1 84-84a84.09 84.09 0 0 1-84 84Zm27.6-88.91a36 36 0 1 0-55.2 0a40 40 0 1 0 55.2 0ZM116 100a12 12 0 1 1 12 12a12 12 0 0 1-12-12Zm12 68a16 16 0 1 1 16-16a16 16 0 0 1-16 16Z"/>`),
		g.Group(children),
	)
}
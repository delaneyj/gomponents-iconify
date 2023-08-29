package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberEightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M162.44 118.91a52 52 0 1 0-68.88 0a60 60 0 1 0 68.88 0ZM100 80a28 28 0 1 1 28 28a28 28 0 0 1-28-28Zm28 124a36 36 0 1 1 36-36a36 36 0 0 1-36 36Z"/>`),
		g.Group(children),
	)
}
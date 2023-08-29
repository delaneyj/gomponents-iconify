package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotchesThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m210.83 130.83l-80 80a4 4 0 1 1-5.66-5.66l80-80a4 4 0 1 1 5.66 5.66Zm-16-93.66a4 4 0 0 0-5.66 0l-152 152a4 4 0 0 0 5.66 5.66l152-152a4 4 0 0 0 0-5.66Z"/>`),
		g.Group(children),
	)
}
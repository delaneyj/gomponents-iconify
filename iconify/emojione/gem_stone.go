package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GemStone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#9450e0" d="M41 4H23L2 20.1L32 60l30-39.9z"/><path fill="#c28fef" d="m32 60l12.5-39.9H18.8zM9.5 9.5L2 20.1h16.8L23 4zm45 0L41 4l3.5 16.1H62z"/>`),
		g.Group(children),
	)
}
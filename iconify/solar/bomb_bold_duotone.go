package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BombBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9.5 22a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15Z" opacity=".5"/><path d="M17.981 2.353a.558.558 0 0 1 1.038 0l.654 1.66c.057.143.17.257.315.314l1.659.654c.47.186.47.852 0 1.038l-1.66.654a.558.558 0 0 0-.314.315l-.654 1.659a.558.558 0 0 1-1.038 0l-.654-1.66a.558.558 0 0 0-.315-.314l-1.659-.654a.558.558 0 0 1 0-1.038l1.66-.654a.558.558 0 0 0 .314-.315l.654-1.659Z"/><path d="m16.477 6.462l-2.23 2.23c.387.318.743.673 1.06 1.061l2.23-2.23l-.21-.536a.558.558 0 0 0-.315-.314l-.535-.211Z" opacity=".7"/></g>`),
		g.Group(children),
	)
}
package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhitePentagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FFF" d="M34.824 10.451a2.002 2.002 0 0 1 2.352 0l25.229 18.33c.701.51.994 1.412.727 2.236l-9.637 29.659a2 2 0 0 1-1.902 1.382H20.407a2 2 0 0 1-1.902-1.382L8.868 31.017a2.002 2.002 0 0 1 .727-2.236l25.229-18.33z"/><path fill="none" stroke="#000" stroke-width="2" d="M34.824 10.451a2.002 2.002 0 0 1 2.352 0l25.229 18.33c.701.51.994 1.412.727 2.236l-9.637 29.659a2 2 0 0 1-1.902 1.382H20.407a2 2 0 0 1-1.902-1.382L8.868 31.017a2.002 2.002 0 0 1 .727-2.236l25.229-18.33z"/>`),
		g.Group(children),
	)
}
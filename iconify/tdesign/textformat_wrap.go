package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextformatWrap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4h20v2H2V4Zm0 7h15c1.148 0 2.38.284 3.35 1.012C21.36 12.77 22 13.946 22 15.5s-.64 2.73-1.65 3.488C19.38 19.716 18.148 20 17 20h-5.414l1.707-1.707l2-2l.707-.707L17.414 17l-.707.707l-.293.293H17c.852 0 1.62-.216 2.15-.613c.49-.367.85-.941.85-1.887c0-.946-.36-1.52-.85-1.887C18.62 13.216 17.852 13 17 13H2v-2Zm1 7H2v2h8v-2H3Z"/>`),
		g.Group(children),
	)
}
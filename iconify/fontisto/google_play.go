package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.751.61l13.124 7.546l-2.813 2.813zM1.032 0l12.047 12L1.033 24a1.72 1.72 0 0 1-1.032-1.643v.003V1.641L0 1.576C0 .875.42.272 1.022.005l.011-.004zm19.922 10.594c.414.307.679.795.679 1.344l-.001.065V12a1.558 1.558 0 0 1-.629 1.403l-.004.003l-2.813 1.593L15.14 12l3.047-3.047zM3.751 23.39l10.312-10.359l2.813 2.813z"/>`),
		g.Group(children),
	)
}
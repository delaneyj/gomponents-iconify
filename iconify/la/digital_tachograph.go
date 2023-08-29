package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitalTachograph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.018 7A2.017 2.017 0 0 0 2 9.018v13.964C2 24.097 2.903 25 4.018 25h22.964A2.017 2.017 0 0 0 29 22.982V9.018A2.017 2.017 0 0 0 26.982 7H4.018zM4 9h23v14H4V9zm2 2v2h10v-2H6zm0 4v2h2v-2H6zm4 0v2h2v-2h-2zm4 0v2h2v-2h-2zm-8 4v2h10v-2H6zm12 0v2h7v-2h-7z"/>`),
		g.Group(children),
	)
}
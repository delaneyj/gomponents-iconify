package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 397V35Q0 14 18 6l210 210L18 426q-18-9-18-29zm295-114L65 415l181-181zm71-92q13 10 13 25t-12 25l-49 28l-54-53l54-53zM65 17l230 132l-49 49z"/>`),
		g.Group(children),
	)
}
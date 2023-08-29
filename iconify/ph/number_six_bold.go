package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSixBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 100a59.21 59.21 0 0 0-7.81.53l26.27-46.64a12 12 0 0 0-20.92-11.78L76 130.13A60 60 0 1 0 128 100Zm0 96a36 36 0 1 1 36-36a36 36 0 0 1-36 36Z"/>`),
		g.Group(children),
	)
}
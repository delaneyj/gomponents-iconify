package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlpinejsIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#77C1D2" d="M199.111 0L256 56.639l-56.889 56.64l-56.889-56.64z"/><path fill="#2D3441" d="m56.889 0l117.938 117.421H61.049L0 56.639z"/>`),
		g.Group(children),
	)
}
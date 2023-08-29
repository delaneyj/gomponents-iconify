package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M437.3 30L202.7 339.3L64 200.7l-64 64L213.3 478L512 94z"/>`),
		g.Group(children),
	)
}
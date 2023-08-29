package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 72v24a8 8 0 0 1-16 0V80h-98.45L79.49 202.81a8 8 0 0 1-15 0l-48-128a8 8 0 1 1 15-5.62L72 177.22l40.51-108A8 8 0 0 1 120 64h112a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadicalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 72v24a4 4 0 0 1-8 0V76H122.77l-47 125.4a4 4 0 0 1-7.49 0l-48-128a4 4 0 0 1 7.49-2.81L72 188.61l44.26-118A4 4 0 0 1 120 68h112a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}
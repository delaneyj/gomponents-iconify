package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gavel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 13l-7.5 7.5c-.83.83-2.17.83-3 0c0 0 0 0 0 0a2.12 2.12 0 0 1 0-3L11 10m5 6l6-6M8 8l6-6M9 7l8 8m4-4l-8-8"/>`),
		g.Group(children),
	)
}
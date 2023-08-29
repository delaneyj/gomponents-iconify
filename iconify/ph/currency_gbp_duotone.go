package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyGbpDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M168 208H60a36 36 0 0 0 36-36V84a44 44 0 0 1 72-33.95Z" opacity=".2"/><path d="M192 208a8 8 0 0 1-8 8H56a8 8 0 0 1 0-16h4a28 28 0 0 0 28-28v-36H56a8 8 0 0 1 0-16h32V84a52 52 0 0 1 85.08-40.12a8 8 0 1 1-10.18 12.34A36 36 0 0 0 104 84v36h32a8 8 0 0 1 0 16h-32v36a43.82 43.82 0 0 1-10.08 28H184a8 8 0 0 1 8 8Z"/></g>`),
		g.Group(children),
	)
}
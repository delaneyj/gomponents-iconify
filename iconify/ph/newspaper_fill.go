package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewspaperFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 48H56a16 16 0 0 0-16 16v120a8 8 0 0 1-16 0V88a8 8 0 0 0-16 0v96.11A24 24 0 0 0 32 208h176a24 24 0 0 0 24-24V64a16 16 0 0 0-16-16Zm-40 104H96a8 8 0 0 1 0-16h80a8 8 0 0 1 0 16Zm0-32H96a8 8 0 0 1 0-16h80a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}
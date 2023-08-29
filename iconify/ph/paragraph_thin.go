package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 44H96a60 60 0 0 0 0 120h44v44a4 4 0 0 0 8 0V52h32v156a4 4 0 0 0 8 0V52h20a4 4 0 0 0 0-8Zm-68 112H96a52 52 0 0 1 0-104h44Z"/>`),
		g.Group(children),
	)
}
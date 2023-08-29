package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 36H96a68 68 0 0 0 0 136h36v36a12 12 0 0 0 24 0V60h16v148a12 12 0 0 0 24 0V60h12a12 12 0 0 0 0-24Zm-76 112H96a44 44 0 0 1 0-88h36Z"/>`),
		g.Group(children),
	)
}
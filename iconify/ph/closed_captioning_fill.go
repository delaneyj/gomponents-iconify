package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptioningFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 48H32a16 16 0 0 0-16 16v128a16 16 0 0 0 16 16h192a16 16 0 0 0 16-16V64a16 16 0 0 0-16-16ZM116 162.64a40 40 0 1 1 0-69.28a8 8 0 1 1-8 13.85a24 24 0 1 0 0 41.58a8 8 0 0 1 8 13.85Zm80 0a40 40 0 1 1 0-69.28a8 8 0 1 1-8 13.85a24 24 0 1 0 0 41.58a8 8 0 0 1 8 13.85Z"/>`),
		g.Group(children),
	)
}
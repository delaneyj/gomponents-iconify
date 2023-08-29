package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatRoundDotsBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10c-1.6 0-3.112-.376-4.452-1.044a1.634 1.634 0 0 0-1.149-.133l-2.226.596a1.3 1.3 0 0 1-1.591-1.592l.595-2.226a1.633 1.633 0 0 0-.134-1.148A9.96 9.96 0 0 1 2 12C2 6.477 6.477 2 12 2s10 4.477 10 10ZM8 13a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm3-1a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm4 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
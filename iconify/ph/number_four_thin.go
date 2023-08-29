package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFourThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M180 152a4 4 0 0 1-4 4h-20v52a4 4 0 0 1-8 0v-52H72a4 4 0 0 1-3.77-5.35l40-112a4 4 0 1 1 7.54 2.69L77.68 148H148V96a4 4 0 0 1 8 0v52h20a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}
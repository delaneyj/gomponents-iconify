package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirtualRealityFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 116a12 12 0 0 1-12 12h-12v-24h12a12 12 0 0 1 12 12Zm80 12a80.09 80.09 0 0 1-80 80H80a80 80 0 0 1 0-160h96a80.09 80.09 0 0 1 80 80ZM119 88.59A8 8 0 0 0 108.59 93L90 138.74L71.41 93a8 8 0 1 0-14.82 6l26 64a8 8 0 0 0 14.82 0l26-64A8 8 0 0 0 119 88.59Zm61.83 49.76A28 28 0 0 0 164 88h-20a8 8 0 0 0-8 8v64a8 8 0 0 0 16 0v-16h12c.54 0 1.08 0 1.61-.05l11.44 20A8 8 0 0 0 191 156Z"/>`),
		g.Group(children),
	)
}
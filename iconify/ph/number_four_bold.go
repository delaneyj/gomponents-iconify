package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFourBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M188 152a12 12 0 0 1-12 12h-12v44a12 12 0 0 1-24 0v-44H72a12 12 0 0 1-11.3-16l40-112a12 12 0 1 1 22.6 8L89 140h51V96a12 12 0 0 1 24 0v44h12a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}
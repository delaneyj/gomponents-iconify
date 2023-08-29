package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberThreeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M180 160a60 60 0 0 1-100 44.72a12 12 0 1 1 16-17.88A36 36 0 1 0 120 124a12 12 0 0 1-9.6-19.2L144 60H88a12 12 0 0 1 0-24h80a12 12 0 0 1 9.6 19.2l-36.48 48.64A60.11 60.11 0 0 1 180 160Z"/>`),
		g.Group(children),
	)
}
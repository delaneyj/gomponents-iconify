package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitInsetFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm37.66 141.66a8 8 0 0 1-11.32 0l-64-64a8 8 0 0 1 11.32-11.32l64 64a8 8 0 0 1 0 11.32Z"/>`),
		g.Group(children),
	)
}
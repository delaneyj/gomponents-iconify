package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowURightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m221.66 181.66l-48 48a8 8 0 0 1-11.32 0l-48-48a8 8 0 0 1 11.32-11.32L160 204.69V88a48 48 0 0 0-96 0v88a8 8 0 0 1-16 0V88a64 64 0 0 1 128 0v116.69l34.34-34.35a8 8 0 0 1 11.32 11.32Z"/>`),
		g.Group(children),
	)
}
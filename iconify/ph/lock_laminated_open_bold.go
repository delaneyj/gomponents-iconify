package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockLaminatedOpenBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 76H100V56a28 28 0 0 1 28-28c13.51 0 25.65 9.62 28.24 22.39a12 12 0 1 0 23.52-4.78C174.87 21.5 153.1 4 128 4a52.06 52.06 0 0 0-52 52v20H48a20 20 0 0 0-20 20v112a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V96a20 20 0 0 0-20-20ZM52 144h152v16H52Zm152-44v20H52v-20ZM52 204v-20h152v20Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 48v104a20 20 0 0 1-20 20h-95.08a12 12 0 0 1-17.41 16.49l-20-20a12 12 0 0 1 0-17l20-20A12 12 0 0 1 112.92 148H204V52H100a12 12 0 0 1-24 0v-4a20 20 0 0 1 20-20h112a20 20 0 0 1 20 20Zm-60 144a12 12 0 0 0-12 12H52v-96h91.08a12 12 0 0 0 17.41 16.49l20-20a12 12 0 0 0 0-17l-20-20A12 12 0 0 0 143.08 84H48a20 20 0 0 0-20 20v104a20 20 0 0 0 20 20h112a20 20 0 0 0 20-20v-4a12 12 0 0 0-12-12Z"/>`),
		g.Group(children),
	)
}
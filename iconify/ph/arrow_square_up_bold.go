package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152ZM87.51 128.49a12 12 0 0 1 0-17l32-32a12 12 0 0 1 17 0l32 32a12 12 0 0 1-17 17L140 117v51a12 12 0 0 1-24 0v-51l-11.51 11.52a12 12 0 0 1-16.98-.03Z"/>`),
		g.Group(children),
	)
}
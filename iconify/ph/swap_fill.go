package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 48v104a16 16 0 0 1-16 16h-96v16a8 8 0 0 1-13.66 5.66l-24-24a8 8 0 0 1 0-11.32l24-24A8 8 0 0 1 112 136v16h96V48H96v8a8 8 0 0 1-16 0v-8a16 16 0 0 1 16-16h112a16 16 0 0 1 16 16Zm-56 144a8 8 0 0 0-8 8v8H48V104h96v16a8 8 0 0 0 13.66 5.66l24-24a8 8 0 0 0 0-11.32l-24-24A8 8 0 0 0 144 72v16H48a16 16 0 0 0-16 16v104a16 16 0 0 0 16 16h112a16 16 0 0 0 16-16v-8a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}
package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Creditcard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 20h8M3 13h26m-6 7h2M5 7C4 7 3 8 3 9v14c0 1 1 2 2 2h22c1 0 2-1 2-2V9c0-1-1-2-2-2H5Z"/>`),
		g.Group(children),
	)
}
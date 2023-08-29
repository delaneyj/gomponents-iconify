package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PollH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5H5zm2 2h18v18H7V7zm3 3v2h8v-2h-8zm0 5v2h12v-2H10zm0 5v2h6v-2h-6z"/>`),
		g.Group(children),
	)
}
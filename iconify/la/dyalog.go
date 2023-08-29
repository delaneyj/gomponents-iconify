package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dyalog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 5v7h2V7h7c4.962 0 9 4.038 9 9s-4.038 9-9 9H6v2h9c6.065 0 11-4.935 11-11S21.065 5 15 5H6z"/>`),
		g.Group(children),
	)
}
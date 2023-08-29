package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseCornerArrowOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 42.7L469.3 0l-128 128V0l-64 64l-.1 170.8l170.8-.1l64-64H384l128-128zM0 341.3h128L0 469.3L42.7 512l128-128v128l64-64l.1-170.8l-170.8.1l-64 64z"/>`),
		g.Group(children),
	)
}
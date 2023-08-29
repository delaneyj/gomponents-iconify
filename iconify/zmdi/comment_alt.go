package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m426 45l1 384l-86-85H43q-18 0-30.5-12.5T0 301V45q0-17 12.5-29.5T43 3h341q18 0 30 12.5T426 45z"/>`),
		g.Group(children),
	)
}
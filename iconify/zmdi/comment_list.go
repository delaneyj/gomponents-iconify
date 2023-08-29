package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H85L0 429V45q0-17 12.5-29.5T43 3h341zM128 259v-43H85v43h43zm0-64v-43H85v43h43zm0-64V88H85v43h43zm149 128v-43H171v43h106zm64-64v-43H171v43h170zm0-64V88H171v43h170z"/>`),
		g.Group(children),
	)
}
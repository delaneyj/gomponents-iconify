package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentAltText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v384l-86-85H43q-18 0-30.5-12.5T0 301V45q0-17 12.5-29.5T43 3h341zm-43 256v-43H85v43h256zm0-64v-43H85v43h256zm0-64V88H85v43h256z"/>`),
		g.Group(children),
	)
}
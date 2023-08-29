package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 3q18 0 30.5 12.5T427 45v256q0 18-12.5 30.5T384 344H85L0 429V45q0-17 12.5-29.5T43 3h341zM64 259h299l-96-128l-75 96l-53-64z"/>`),
		g.Group(children),
	)
}
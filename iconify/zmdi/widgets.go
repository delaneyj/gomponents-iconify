package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Widgets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213 245h171v171H213V245zM0 416V245h171v171H0zM0 32h171v171H0V32zM291 4l121 121l-121 120l-120-120z"/>`),
		g.Group(children),
	)
}
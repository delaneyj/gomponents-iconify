package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Navicon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M96 241h320v32H96z" fill="currentColor"/><path d="M96 145h320v32H96z" fill="currentColor"/><path d="M96 337h320v32H96z" fill="currentColor"/>`),
		g.Group(children),
	)
}
package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HappyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zm3 4c.552 0 1 .672 1 1.5S11.552 7 11 7s-1-.672-1-1.5s.448-1.5 1-1.5zM5 4c.552 0 1 .672 1 1.5S5.552 7 5 7s-1-.672-1-1.5S4.448 4 5 4zm3 10c-2.607 0-4.772-2.186-5-4.973c1.465.846 3.188 1.329 5 1.329s3.535-.481 5-1.327C12.772 11.817 10.607 14 8 14z"/>`),
		g.Group(children),
	)
}
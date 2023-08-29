package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DzFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" d="M320 0h320v480H320z"/><path fill="#006233" d="M0 0h320v480H0z"/><path fill="#d21034" d="M424 180a120 120 0 1 0 0 120a96 96 0 1 1 0-120m4 60l-108-35.2l67.2 92V183.2l-67.2 92z"/>`),
		g.Group(children),
	)
}
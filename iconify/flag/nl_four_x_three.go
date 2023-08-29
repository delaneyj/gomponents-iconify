package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NlFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#21468b" d="M0 0h640v480H0z"/><path fill="#fff" d="M0 0h640v320H0z"/><path fill="#ae1c28" d="M0 0h640v160H0z"/>`),
		g.Group(children),
	)
}
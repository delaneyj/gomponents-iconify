package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#D91023" d="M0 0h640v480H0z"/><path fill="#fff" d="M213.3 0h213.4v480H213.3z"/>`),
		g.Group(children),
	)
}
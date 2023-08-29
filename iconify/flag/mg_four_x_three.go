package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MgFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#fc3d32" d="M213.3 0H640v240H213.3z"/><path fill="#007e3a" d="M213.3 240H640v240H213.3z"/><path fill="#fff" d="M0 0h213.3v480H0z"/></g>`),
		g.Group(children),
	)
}
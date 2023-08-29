package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CiFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#00cd00" d="M426.8 0H640v480H426.8z"/><path fill="#ff9a00" d="M0 0h212.9v480H0z"/><path fill="#fff" d="M212.9 0h214v480h-214z"/></g>`),
		g.Group(children),
	)
}
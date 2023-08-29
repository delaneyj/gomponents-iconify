package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MlFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="red" d="M425.8 0H640v480H425.7z"/><path fill="#009a00" d="M0 0h212.9v480H0z"/><path fill="#ff0" d="M212.9 0h214v480h-214z"/></g>`),
		g.Group(children),
	)
}
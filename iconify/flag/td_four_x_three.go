package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TdFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#002664" d="M0 0h214v480H0z"/><path fill="#c60c30" d="M426 0h214v480H426z"/><path fill="#fecb00" d="M214 0h212v480H214z"/></g>`),
		g.Group(children),
	)
}
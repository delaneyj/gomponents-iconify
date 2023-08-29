package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GaFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#ffe700" d="M640 480H0V0h640z"/><path fill="#36a100" d="M640 160H0V0h640z"/><path fill="#006dbc" d="M640 480H0V320h640z"/></g>`),
		g.Group(children),
	)
}
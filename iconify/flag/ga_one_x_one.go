package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GaOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#ffe700" d="M512 512H0V0h512z"/><path fill="#36a100" d="M512 170.7H0V0h512z"/><path fill="#006dbc" d="M512 512H0V341.3h512z"/></g>`),
		g.Group(children),
	)
}
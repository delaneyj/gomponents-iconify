package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NdbNeos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.293 8.052L5.5 23.361V8.052h26.793zm6.379 6.379L11.88 42.5h26.792V14.431zM5.5 28.465V42.5l37-37l-37 22.965z"/>`),
		g.Group(children),
	)
}
package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConstantSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 5.8A.8.8 0 0 1 .8 5h14.4a.8.8 0 0 1 0 1.6H.8a.8.8 0 0 1-.8-.8Zm.8 4.8a.8.8 0 0 1 .8.8v1.333h3.2V11.4a.8.8 0 0 1 1.6 0v1.333h3.2V11.4a.8.8 0 0 1 1.6 0v1.333h3.2V11.4a.8.8 0 0 1 1.6 0v2.933H0V11.4a.8.8 0 0 1 .8-.8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
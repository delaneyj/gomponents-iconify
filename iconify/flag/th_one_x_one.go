package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#f4f5f8" d="M0 0h512v512H0z"/><path fill="#2d2a4a" d="M0 173.4h512V344H0z"/><path fill="#a51931" d="M0 0h512v88H0zm0 426.7h512V512H0z"/></g>`),
		g.Group(children),
	)
}
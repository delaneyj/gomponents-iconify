package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#f4f5f8" d="M0 0h640v480H0z"/><path fill="#2d2a4a" d="M0 162.5h640v160H0z"/><path fill="#a51931" d="M0 0h640v82.5H0zm0 400h640v80H0z"/></g>`),
		g.Group(children),
	)
}
package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#c10000" d="M0 0h640v480H0z"/><path fill="#fff" d="M0 0h250v200.3H0z"/><g fill="#c10000"><path d="M102.8 31.2h39.9v139.6h-39.8z"/><path d="M192.6 81v40H53V81z"/></g></g>`),
		g.Group(children),
	)
}
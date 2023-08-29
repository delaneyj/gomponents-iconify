package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#ffe800" d="M0 0h640v480H0z"/><path fill="#00148e" d="M0 240h640v240H0z"/><path fill="#da0010" d="M0 360h640v120H0z"/></g>`),
		g.Group(children),
	)
}
package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#fff" d="M640 480H0V0h640z"/><path fill="#388d00" d="M640 480H0V320h640z"/><path fill="#d43516" d="M640 160.1H0V.1h640z"/></g>`),
		g.Group(children),
	)
}
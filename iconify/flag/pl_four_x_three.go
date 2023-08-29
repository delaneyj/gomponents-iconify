package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#fff" d="M640 480H0V0h640z"/><path fill="#dc143c" d="M640 480H0V240h640z"/></g>`),
		g.Group(children),
	)
}
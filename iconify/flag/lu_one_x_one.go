package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#00a1de" d="M0 256h512v256H0z"/><path fill="#ed2939" d="M0 0h512v256H0z"/><path fill="#fff" d="M0 170.7h512v170.6H0z"/>`),
		g.Group(children),
	)
}
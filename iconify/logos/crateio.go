package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crateio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#009DC7" d="M192 64V0h-64v64H0v64h64v64h64v-64h128V64z"/>`),
		g.Group(children),
	)
}
package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Functions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 21v64H107l106 107l-106 107h149v64H0v-43l139-128L0 64V21h256z"/>`),
		g.Group(children),
	)
}
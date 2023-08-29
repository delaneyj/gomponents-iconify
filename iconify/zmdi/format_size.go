package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatSize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M128 21h277v64H299v256h-64V85H128V21zM0 192v-64h192v64h-64v149H64V192H0z"/>`),
		g.Group(children),
	)
}
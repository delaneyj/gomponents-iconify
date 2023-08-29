package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatStrikethrough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M149 341v-64h86v64h-86zM43 21h298v64H235v64h-86V85H43V21zM0 235v-43h384v43H0z"/>`),
		g.Group(children),
	)
}
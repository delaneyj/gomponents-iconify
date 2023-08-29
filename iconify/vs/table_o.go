package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M64 606L0 576V448L832 64l832 384v128l-64 30v534q0 38-26.5 57t-69.5 19t-69.5-19t-26.5-57V694L928 916v480q0 38-26.5 57t-69.5 19t-69.5-19t-26.5-57V916L256 694v446q0 38-26.5 57t-69.5 19t-69.5-19t-26.5-57V606zm768 226l768-352l-768-352L64 480z"/>`),
		g.Group(children),
	)
}
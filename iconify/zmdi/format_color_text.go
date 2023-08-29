package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatColorText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 363h512v85H0v-85zM235 0h42l117 299h-48l-23-64H189l-24 64h-48zm-30 192h102L256 57z"/>`),
		g.Group(children),
	)
}
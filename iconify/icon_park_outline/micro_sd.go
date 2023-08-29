package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicroSd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 7a2 2 0 0 1 2-2h22a2 2 0 0 1 2 2v34a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2v-5l5-4v-4H8v-3l5-6V7Zm19 4v4m-12-4v4m6-4v4"/>`),
		g.Group(children),
	)
}
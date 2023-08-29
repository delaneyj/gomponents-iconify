package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 36.333h2.056v2.056H9.61v2.056h2.056V42.5h4.11v-2.056h2.057V38.39h2.055v-2.056h2.056v-2.055h4.11v2.055h2.056v2.056h2.056v2.056h2.056V42.5h4.11v-2.056h2.056V38.39h2.056v-2.056H42.5v-4.111h-2.055v-2.056h-2.056v-2.055h-2.055v-2.056h-2.056v-4.111h2.056v-2.056h2.055v-2.055h2.057v-2.056H42.5v-4.111h-2.055V9.61h-2.056V7.554h-2.055V5.5h-4.111v2.056h-2.056V9.61h-2.056v2.056h-2.055v2.056h-4.111v-2.056h-2.056V9.61h-2.055V7.554h-2.056V5.5h-4.111v2.056H9.61V9.61H7.556v2.056H5.5v4.111h2.056v2.056H9.61v2.055h2.056v2.056h2.055v4.111h-2.055v2.056H9.61v2.055H7.555v2.056H5.5l-.001 4.11Z"/>`),
		g.Group(children),
	)
}
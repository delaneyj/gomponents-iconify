package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvBrowserUpdateplugin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.207 24.21h-6.091l-5.167 12.705l-7.898-25.831l-5.167 12.705h-6.09"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.576 11.573L36.19 14.29l3.016 1.992l-3.607.235l.724 3.541l-2.717-2.385l-1.992 3.017l-.235-3.608l-3.542.724l2.385-2.716l-3.016-1.993l3.607-.235l-.724-3.541l2.717 2.385l1.992-3.016l.235 3.607Z"/>`),
		g.Group(children),
	)
}
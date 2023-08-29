package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShortcutToUrl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.5 5.5h-29l-4 4v29l4 4h29l4-4v-29l-4-4Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.848 38.896c-7.728-4.517-1.827-9.443 7.115-15.185l6.778 6.433V11.538H18.135l5.93 5.93c-15.688 14.79-10.352 18.36-1.217 21.427Z"/>`),
		g.Group(children),
	)
}
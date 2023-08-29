package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oneplus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.222 20.109V6.5m-7.331 6.278H44.5m-6.278 11.876V41.5H9.5V12.778h16.846m-1.838 21.816V20.109"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.508 20.109h0a4.192 4.192 0 0 1-4.191 4.192h0"/>`),
		g.Group(children),
	)
}
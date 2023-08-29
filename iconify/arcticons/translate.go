package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Translate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.692 44.432L17.308 3.568m5.121 33.069l-5.656-17.402L10.9 36.637m1.958-5.873h7.613m6.765-19.605c-.032 4.543-.26 11.617 5.16 17.328"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.924 13.995a38.588 38.588 0 0 0 15.41-1.733M23.785 23.345c2.586-4.394 13.61-7.906 14.55-.731a5.532 5.532 0 0 1-2.188 5.17"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.446 16.059c.169 5.193-6.115 13.58-10.115 12.011"/>`),
		g.Group(children),
	)
}